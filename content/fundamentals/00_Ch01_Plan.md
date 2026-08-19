# Chapter 1 — Go Fundamentals (~25 min)

Format per lesson: **one concept slide → three tasks** (Beginner → Intermediate → Advanced). No lesson introduces more than one new idea.

| # | Lesson | Time |
| --- | -------- | ------ |
| 1.1 | Declaring Variables with `var` | 3 min |
| 1.2 | Short Declaration `:=` | 2 min |
| 1.3 | Zero Values | 2 min |
| 1.4 | Basic Types Tour | 3 min |
| 1.5 | Constants | 2 min |
| 1.6 | `iota` for Enums | 3 min |
| 1.7 | Operators | 3 min |
| 1.8 | Type Conversion | 3 min |
| 1.9 | `fmt` Print Verbs | 2 min |
| 1.10 | String Formatting Practice | 2 min |
| 1.11 | Getting User Input | 3 min |
| — | Chapter 1 Quiz (5 Qs) | 2 min |
| | **Total** | **~30 min** |

---

### Lesson 1.1 — Declaring Variables with `var`

**Concept slide**

- `var <name> <type> = <value>`, type inference, multi-declaration.
- `var` works at package level; `:=` does not.

**Tasks**
- Beginner: `city` string via `var`, print it.
- Intermediate: `a, b int = 12, 5` on one line, multiply and print.
- Advanced: package-level `language` + `var` block with `year`/`isPublic`.

---

### Lesson 1.2 — Short Declaration `:=`

**Concept slide**

- `name := value` infers the type; equivalent to `var name = value`.
- Only inside functions; needs at least one new variable.

**Tasks**
- Beginner: `city` via `:=`, print it.
- Intermediate: `a, b, c := 10, 20, 30` on one line, print sum.
- Advanced: swap `x, y = y, x` with multiple assignment (no temp).

---

### Lesson 1.3 — Zero Values

**Concept slide**

- Uninitialized variables get zero values: `0`, `0.0`, `""`, `false`, `nil`.
- Declare-then-assign split.

**Tasks**
- Beginner: `score int` / `passed bool` zero values, print `0 false`.
- Intermediate: `count`/`message`/`passed` zero values, bracket-wrapped string.
- Advanced: four vars, set `count = 9` and `name = "Go"`, keep others at zero.

---

### Lesson 1.4 — Basic Types Tour

**Concept slide**

- `int`, `float64`, `string`, `bool`; `fmt.Println` infers conversion.

**Tasks**
- Beginner: one var of each type, print on one line.
- Intermediate: `:=` + an expression using each type.
- Advanced: combine types — string concat, comparison → bool, float.

---

### Lesson 1.5 — Constants

**Concept slide**

- `const <name> = <value>`; compile-time values; grouped `const` blocks.

**Tasks**
- Beginner: `const MaxScore = 100`, print it.
- Intermediate: `const` block with `Pi`, `E`, `Phi`.
- Advanced: `Radius = 7.0` + `Pi` constant, compute circumference.

---

### Lesson 1.6 — `iota` for Enums

**Concept slide**

- `const` block auto-counter `iota` (0, 1, 2, …); skip with `_`; `1 << iota`.

**Tasks**
- Beginner: `Low`, `Medium`, `High`; `PriorityValue()` returns `Medium` (`1`).
- Intermediate: `Red`, `Green`, `Blue`, `Alpha` skipping `0` with `_`.
- Advanced: flag constants as powers of two via `1 << iota`.

---

### Lesson 1.7 — Operators

**Concept slide**

- Arithmetic `+ - * / %`, comparisons, `&&` / `||` / `!`.

**Tasks**
- Beginner: `IsEven(n int) bool` with `%`.
- Intermediate: `IsLeapYear(year int) bool` — `%`, `&&`, `||`.
- Advanced: `IsBetween(n, low, high int) bool` with `&&`.

---

### Lesson 1.8 — Type Conversion

**Concept slide**

- Explicit casts only: `float64(x)`, `int(y)`; int division truncates.

**Tasks**
- Beginner: `Average(a, b int) float64` — convert before dividing.
- Intermediate: `Celsius(f float64) float64` — Fahrenheit → Celsius.
- Advanced: `ToPixels(inches, dpi float64) int` — float multiply, then truncating `int(...)`.

---

### Lesson 1.9 — `fmt` Print Verbs

**Concept slide**

- `%d`, `%s`, `%q`, `%T`, `%f`/`%.2f`, width/padding flags.

**Tasks**
- Beginner: `%d` + `%T` on `age := 30` → `Age: 30, Type: int`.
- Intermediate: `%s`, `%q`, `%T` on `name := "Gopher"`.
- Advanced: `%04d` zero-padding + `%.2f` rounding on `score`/`pi`.

---

### Lesson 1.10 — String Formatting Practice

**Concept slide**

- `fmt.Sprintf` builds (not prints) strings; multi-line via `\n`; width flags.

**Tasks**
- Beginner: `Receipt(item, qty, price)` → single-line formatted receipt.
- Intermediate: `Report(name, orders, revenue)` → multi-line `\n` report.
- Advanced: `Row(item, price)` → right/left-aligned table row with `%-12s` + `%10.2f`.

---

### Lesson 1.11 — Getting User Input

**Concept slide**

- `fmt.Scanln(&a, &b)` reads space-separated values until a newline; `fmt.Scan(&a, &b)` reads until whitespace.
- Pointers (`&name`) are required so the function can write values back into variables.

**Tasks**
- Beginner: `Greet()` → `fmt.Scanln` a name, return `"Hello, <name>!"`.
- Intermediate: `Sum()` → `fmt.Scan` two ints, return their sum.
- Advanced: `AverageUntilZero()` → loop `fmt.Scan` ints until `0`, return average.

---

## 🧩 Chapter 1 Quiz (5 questions, ~2 min)

1. Which two forms of variable declaration did you learn first? *(`var` and `:=`)*
2. What's the zero value of a `string`? *(`""`)*
3. What keyword declares a compile-time constant? *(`const`)*
4. True/False: Go auto-converts `int` to `float64` when needed. *(False — conversions must be explicit)*
5. Which `fmt` verb prints the type of a value? *(`%T`)*