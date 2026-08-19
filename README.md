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
| **`[← ↑ ↓ →]`** | Navigate within the focused panel (chapters, lessons, task grid) |
| **`[Tab]` / `[Shift+Tab]`** | Switch between the sidebar, task grid and details panels |
| **`[r]` or `[Enter]`** | Run unit tests for selected exercise (with running fox ASCII animation) |
| **`[h]`** | Toggle formatted Markdown task explanation & hints |
| **`[u]` / `[m]`** | Mark / unmark the selected task as completed (lessons and chapters show a green tick once complete) |
| **`[o]`** | Open the selected task in your IDE (asks once on first use, then remembers) |
| **`[s]`** | Re-open the "Open With" picker to switch the IDE |
| **`[n]` / `[p]`** | Jump to Next / Previous exercise within the current lesson |
| **`[q]` / `[Esc]`** | Quit Gostlings |

