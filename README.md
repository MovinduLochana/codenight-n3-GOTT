# Gostlings — Interactive Go Exercises 🦊

An interactive, terminal-based Rustlings-style exercise runner for learning **Go (Golang)**, developed by the **Mozilla Campus Club of SLIIT**.

---

## 🚀 Quick Start & Installation

Students can download the CLI and exercise files onto their machine with a single terminal command:

### 🪟 Windows (PowerShell)
Open PowerShell in your desired workspace directory and run:

```powershell
iwr -useb https://raw.githubusercontent.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/main/install.ps1 | iex
```

Then launch the runner:
```powershell
.\gostlings.exe
```

---

### 🍎 macOS & 🐧 Linux (Terminal)
Open Terminal in your desired workspace directory and run:

```bash
curl -fsSL https://raw.githubusercontent.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/main/install.sh | bash
```

Then launch the runner:
```bash
./gostlings
```

---

## 🎮 TUI Controls & Shortcuts

| Hotkey | Action |
| :--- | :--- |
| **`[← ↑ ↓ →]` / `[h j k l]`** | Navigate across exercise cards in the grid |
| **`[r]` or `[Enter]`** | Run unit tests for selected exercise (with running fox ASCII animation) |
| **`[h]`** | Toggle formatted Markdown task explanation & hints |
| **`[n]` / `[p]`** | Jump to Next / Previous exercise |
| **`[q]` / `[Esc]`** | Quit Gostlings |

---

## 📂 Repository Structure

```text
├── cli/                  # Go TUI Runner source code (Bubble Tea + Lip Gloss)
├── content/              # Lecture content & exercise markdown task explanations
├── exercises/            # Go exercise starter files & companion unit tests
├── web/                  # Next.js 16 web application
├── install.ps1           # Windows PowerShell installer script
├── install.sh            # macOS/Linux Bash installer script
└── .github/workflows/    # CI/CD GitHub Actions release pipeline
```

---

## 🌐 Web Application

The repository also includes a Next.js 16 web app located in `/web` providing lecture slide content, task breakdowns, and evaluated MCQs.

To run the web app locally:
```bash
cd web
npm install
npm run dev
```
