# Go Programming Assessment (Chapters 1–5)

This directory contains a set of 10 assessment questions designed to evaluate your understanding of the concepts covered in the first 5 chapters:
1. **Go Fundamentals** (Variables, short declarations, zero values, types, constants, iota, operators, type conversion, formatting)
2. **Control Flow** (`if/else`, loops, `switch`)
3. **Collections** (Arrays, slices, maps, `range`)
4. **Functions** (Multiple return values, variadic functions, closures, `defer`)
5. **Structs & Interfaces** (Struct definition, pointers, methods, interfaces, embedding)

---

## 📋 Assessment Index

| Q# | Level | Topic | Key Concepts | Directory |
| :--- | :--- | :--- | :--- | :--- |
| **01** | 🟢 Beginner | **Format & Conversion** | Variable declaration, explicit type casting, `fmt.Sprintf` | `[01_Beginner_Formatting](./01_Beginner_Formatting)` |
| **02** | 🟢 Beginner | **FizzBuzz Loop** | `for` loops, `if`/`else` branching, modulo operator | `[02_Beginner_FizzBuzz](./02_Beginner_FizzBuzz)` |
| **03** | 🟢 Beginner | **Slices & Append** | Slice initialization, `append()`, `len` and `cap` checks | `[03_Beginner_Slices](./03_Beginner_Slices)` |
| **04** | 🟢 Beginner | **Error Returns** | Multiple return values, error handling, `errors.New` | `[04_Beginner_Functions](./04_Beginner_Functions)` |
| **05** | 🟡 Intermediate | **Iota & Switch Enums** | `iota` enum patterns, custom types, `switch` statement | `[05_Intermediate_Enums](./05_Intermediate_Enums)` |
| **06** | 🟡 Intermediate | **Map Frequency** | Maps, reading/writing keys, `range` loops | `[06_Intermediate_Maps](./06_Intermediate_Maps)` |
| **07** | 🟡 Intermediate | **Structs & Methods** | Struct declaration, pointer vs. value receiver methods | `[07_Intermediate_Structs](./07_Intermediate_Structs)` |
| **08** | 🔴 Advanced | **Filter Closures** | Functions returning functions, closure state capture | `[08_Advanced_Closures](./08_Advanced_Closures)` |
| **09** | 🔴 Advanced | **Interfaces & Embedding** | Struct composition, implicit interfaces, polymorphism | `[09_Advanced_Interfaces](./09_Advanced_Interfaces)` |
| **10** | 🔴 Advanced | **Defer Stack & Pointers** | Pointer mutation, custom stack logic, `defer` LIFO order | `[10_Advanced_Defer_Stack](./10_Advanced_Defer_Stack)` |
| **11** | 🟢 Beginner | **Temp Conversion** | Formula arithmetic, float formatting | `[11_Beginner_Temp_Converter](./11_Beginner_Temp_Converter)` |
| **12** | 🟢 Beginner | **Palindrome Check** | Strings, runes, case insensitivity, loops | `[12_Beginner_Palindrome](./12_Beginner_Palindrome)` |
| **13** | 🟢 Beginner | **Slice Min/Max** | Loops, conditionals, error handling | `[13_Beginner_Min_Max](./13_Beginner_Min_Max)` |
| **14** | 🟢 Beginner | **Pointer Swap** | Pointer referencing and dereferencing | `[14_Beginner_Pointer_Swap](./14_Beginner_Pointer_Swap)` |
| **15** | 🟡 Intermediate | **JSON Struct Tags** | Struct layout, struct tags, JSON marshalling | `[15_Intermediate_JSON](./15_Intermediate_JSON)` |
| **16** | 🟡 Intermediate | **File Line Reader** | File I/O, `bufio.Scanner`, deferred close | `[16_Intermediate_File_Reader](./16_Intermediate_File_Reader)` |
| **17** | 🟢 Beginner | **User Greeting** | `fmt.Scanln`, reading user input, `fmt.Printf` | `[19_Beginner_User_Greeting](./19_Beginner_User_Greeting)` |
| **18** | 🟡 Intermediate | **User Average** | `fmt.Scan`, input loops, accumulation | `[20_Intermediate_User_Average](./20_Intermediate_User_Average)` |

---

## 🚀 How to Complete the Assessment

1. Navigate to the task directory (e.g., `cd 01_Beginner_Formatting`).
2. Read the instructions in `task.md`.
3. Open `main.go` and implement the logic inside the marked `TODO` placeholders.
4. Run your code using:
   ```bash
   go run main.go
   ```
5. Verify that your console output matches the **Expected Output** documented in `task.md`.
