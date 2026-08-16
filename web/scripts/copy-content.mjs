/**
 * Copy ../content → .data/content
 *
 * Runs before dev/build so all filesystem reads are scoped inside the
 * project directory. This prevents Turbopack from tracing the whole
 * monorepo into the output bundle.
 */

import { cpSync, existsSync, mkdirSync, rmSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const projectRoot = join(__dirname, "..");
const source = join(projectRoot, "..", "content");
const dest = join(projectRoot, ".data", "content");

if (!existsSync(source)) {
  console.error("❌ content/ directory not found at", source);
  process.exit(1);
}

// Clean and re-copy to pick up any content changes
if (existsSync(dest)) {
  rmSync(dest, { recursive: true });
}

mkdirSync(dirname(dest), { recursive: true });
cpSync(source, dest, { recursive: true });

console.log("✓ Copied content → .data/content");
