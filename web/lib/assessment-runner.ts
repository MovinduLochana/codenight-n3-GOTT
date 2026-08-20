import "server-only";

const JUDGE0_DEFAULT_URL = "https://judge0.hefttech.com";
const JUDGE0_GO_LANGUAGE_ID = 60; // Go 1.13.5
const RUN_TIMEOUT_MS = 15_000;

export type FixtureFile = { name: string; content: string };

/**
 * Executes Go code and test suite via the Judge0 sandbox API.
 */
export async function runGoTest(
  userCode: string,
  testCode: string,
  fixtures: FixtureFile[] = [],
): Promise<{ passed: boolean; output: string }> {
  try {
    const mergedCode = mergeGoFiles(userCode, testCode);

    const judge0Url = (process.env.JUDGE0_URL || JUDGE0_DEFAULT_URL).replace(
      /\/+$/,
      "",
    );
    const authToken = process.env.JUDGE0_AUTH_TOKEN;

    const headers: Record<string, string> = {
      "Content-Type": "application/json",
    };
    if (authToken) {
      headers["X-Auth-Token"] = authToken;
    }

    const payload: {
      language_id: number;
      source_code: string;
      cpu_time_limit: number;
      additional_files?: string;
    } = {
      language_id: JUDGE0_GO_LANGUAGE_ID,
      source_code: Buffer.from(mergedCode, "utf-8").toString("base64"),
      cpu_time_limit: 10,
    };

    if (fixtures.length > 0) {
      const zipBuffer = createZip(fixtures);
      payload.additional_files = zipBuffer.toString("base64");
    }

    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), RUN_TIMEOUT_MS);

    const response = await fetch(
      `${judge0Url}/submissions?wait=true&base64_encoded=true`,
      {
        method: "POST",
        headers,
        body: JSON.stringify(payload),
        signal: controller.signal,
      },
    );

    clearTimeout(timeoutId);

    if (!response.ok) {
      const errorText = await response.text().catch(() => "");
      return {
        passed: false,
        output: `Judge0 execution request failed (${response.status}): ${errorText || response.statusText}`,
      };
    }

    const data = (await response.json()) as {
      status?: { id: number; description: string };
      stdout?: string | null;
      stderr?: string | null;
      compile_output?: string | null;
      message?: string | null;
    };

    const decode = (str?: string | null): string =>
      str ? Buffer.from(str, "base64").toString("utf-8").trim() : "";

    const stdout = decode(data.stdout);
    const stderr = decode(data.stderr);
    const compileOutput = decode(data.compile_output);
    const message = decode(data.message);
    const statusId = data.status?.id;

    // Status 3: Accepted (Code ran and exited with 0)
    if (statusId === 3) {
      return {
        passed: true,
        output: stdout || "PASS",
      };
    }

    // Status 6: Compilation Error
    if (statusId === 6) {
      return {
        passed: false,
        output: compileOutput || stderr || "Compilation error occurred.",
      };
    }

    // Status 5: Time Limit Exceeded
    if (statusId === 5) {
      return {
        passed: false,
        output: "Test execution timed out (10s limit exceeded).",
      };
    }

    // Status 11: Runtime Error (NZEC - Non-Zero Exit Code, e.g. failed assertions in tests)
    if (statusId === 11) {
      const failureOutput = stdout || stderr || message || compileOutput;
      return {
        passed: false,
        output: failureOutput || "Tests failed.",
      };
    }

    // Other failure states
    const output =
      stderr ||
      stdout ||
      compileOutput ||
      message ||
      data.status?.description ||
      "Execution error.";

    return {
      passed: false,
      output,
    };
  } catch (error) {
    const err = error as Error;
    if (err.name === "AbortError") {
      return {
        passed: false,
        output: "Test execution timed out while communicating with sandbox.",
      };
    }
    return {
      passed: false,
      output: `Failed to execute tests: ${err.message || "Unknown error"}`,
    };
  }
}

/**
 * Merges user Go code with test definitions and creates a testing.Main runner.
 */
function mergeGoFiles(userCode: string, testCode: string): string {
  const importRegex = /import\s*\(([\s\S]*?)\)|import\s+([^\n;]+)/g;
  const imports = new Set<string>();

  function collectImports(code: string) {
    for (const match of code.matchAll(importRegex)) {
      if (match[1]) {
        const lines = match[1]
          .split("\n")
          .map((l) => l.trim())
          .filter(Boolean);
        for (const line of lines) imports.add(line);
      } else if (match[2]) {
        imports.add(match[2].trim());
      }
    }
  }

  collectImports(userCode);
  collectImports(testCode);
  imports.add('"testing"');

  function cleanCode(code: string): string {
    return code
      .replace(/^\s*package\s+\w+[\s;]*/m, "")
      .replace(/import\s*\(([\s\S]*?)\)|import\s+([^\n;]+)/g, "")
      .trim();
  }

  let cleanedUserCode = cleanCode(userCode);
  let cleanedTestCode = cleanCode(testCode);

  // Extract all test functions: func Test*(t *testing.T)
  const testFuncRegex = /func\s+(Test\w+)\s*\(\s*\w+\s+\*testing\.T\s*\)/g;
  const testNames: string[] = [];
  for (const match of cleanedTestCode.matchAll(testFuncRegex)) {
    testNames.push(match[1]);
  }

  // If user code defines `main()`, rename it to `userMain()` to avoid `main redeclared`
  const hasUserMain = /\bfunc\s+main\s*\(\s*\)/.test(cleanedUserCode);
  if (hasUserMain) {
    cleanedUserCode = cleanedUserCode.replace(
      /\bfunc\s+main\s*\(\s*\)/g,
      "func userMain()",
    );
    cleanedTestCode = cleanedTestCode.replace(
      /\bmain\s*\(\s*\)/g,
      "userMain()",
    );
  }

  const runnerMain = `
func main() {
    matchAll := func(pat, str string) (bool, error) { return true, nil }
    tests := []testing.InternalTest{
        ${testNames.map((name) => `{"${name}", ${name}},`).join("\n        ")}
    }
    benchmarks := []testing.InternalBenchmark{}
    examples := []testing.InternalExample{}

    testing.Main(matchAll, tests, benchmarks, examples)
}
`;

  return `package main

import (
    ${Array.from(imports).join("\n    ")}
)

// --- USER CODE ---
${cleanedUserCode}

// --- TEST SUITE ---
${cleanedTestCode}

// --- TEST RUNNER ---
${runnerMain}
`;
}

/**
 * Creates an uncompressed in-memory ZIP buffer for Judge0 additional_files.
 */
function createZip(files: FixtureFile[]): Buffer {
  let offset = 0;
  const encoder = new TextEncoder();
  const localHeaders: Uint8Array[] = [];
  const centralHeaders: Uint8Array[] = [];

  const crc32Table = getCrc32Table();

  for (const file of files) {
    const data = encoder.encode(file.content);
    const nameBytes = encoder.encode(file.name);

    let crc = 0 ^ -1;
    for (let i = 0; i < data.length; i++) {
      crc = (crc >>> 8) ^ crc32Table[(crc ^ data[i]) & 0xff];
    }
    crc = (crc ^ -1) >>> 0;

    // Local file header (30 bytes + name + data)
    const localHeader = new Uint8Array(30 + nameBytes.length + data.length);
    const view = new DataView(localHeader.buffer);
    view.setUint32(0, 0x04034b50, true);
    view.setUint16(4, 20, true);
    view.setUint16(6, 0, true);
    view.setUint16(8, 0, true); // Stored (no compression)
    view.setUint32(14, crc, true);
    view.setUint32(18, data.length, true);
    view.setUint32(22, data.length, true);
    view.setUint16(26, nameBytes.length, true);
    view.setUint16(28, 0, true);
    localHeader.set(nameBytes, 30);
    localHeader.set(data, 30 + nameBytes.length);
    localHeaders.push(localHeader);

    // Central directory header (46 bytes + name)
    const centralHeader = new Uint8Array(46 + nameBytes.length);
    const cView = new DataView(centralHeader.buffer);
    cView.setUint32(0, 0x02014b50, true);
    cView.setUint16(4, 20, true);
    cView.setUint16(6, 20, true);
    cView.setUint16(8, 0, true);
    cView.setUint16(10, 0, true);
    cView.setUint32(16, crc, true);
    cView.setUint32(20, data.length, true);
    cView.setUint32(24, data.length, true);
    cView.setUint16(28, nameBytes.length, true);
    cView.setUint32(42, offset, true);
    centralHeader.set(nameBytes, 46);
    centralHeaders.push(centralHeader);

    offset += localHeader.length;
  }

  const centralDirSize = centralHeaders.reduce((acc, h) => acc + h.length, 0);
  const eocd = new Uint8Array(22);
  const eocdView = new DataView(eocd.buffer);
  eocdView.setUint32(0, 0x06054b50, true);
  eocdView.setUint16(8, files.length, true);
  eocdView.setUint16(10, files.length, true);
  eocdView.setUint32(12, centralDirSize, true);
  eocdView.setUint32(16, offset, true);

  const totalLength = offset + centralDirSize + 22;
  const result = new Uint8Array(totalLength);
  let pos = 0;
  for (const h of localHeaders) {
    result.set(h, pos);
    pos += h.length;
  }
  for (const h of centralHeaders) {
    result.set(h, pos);
    pos += h.length;
  }
  result.set(eocd, pos);

  return Buffer.from(result);
}

let cachedCrc32Table: Uint32Array | null = null;
function getCrc32Table(): Uint32Array {
  if (!cachedCrc32Table) {
    cachedCrc32Table = new Uint32Array(256);
    for (let i = 0; i < 256; i++) {
      let c = i;
      for (let j = 0; j < 8; j++) {
        c = c & 1 ? 0xedb88320 ^ (c >>> 1) : c >>> 1;
      }
      cachedCrc32Table[i] = c;
    }
  }
  return cachedCrc32Table;
}
