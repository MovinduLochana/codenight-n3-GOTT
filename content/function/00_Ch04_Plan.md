# Chapter 4 — Functions (~20 min)

Format per lesson: **one concept slide → three tasks** (Beginner → Intermediate → Advanced).

| # | Lesson | Time |
| --- | -------- | ------ |
| 4.1 | Functions & Multiple Return Values | 5 min |
| 4.2 | Variadic Functions | 5 min |
| 4.3 | Closures | 5 min |
| 4.4 | `defer` | 5 min |
| — | Chapter 4 Quiz (5 Qs) | 2 min |
| | **Total** | **~22 min** |

---

### Lesson 4.1 — Functions & Multiple Return Values

**Concept slide**

- `(value, error)` pattern; multiple returns; caller checks error immediately.
- Note: `error` is covered in depth in Chapter 6.

**Tasks**
- Beginner: `Divide(a, b int) (int, error)` — `errors.New` on `b == 0`.
- Intermediate: `DivMod(a, b int) (q, r int, err error)`.
- Advanced: `MinMax(nums []int) (min, max int)`.

---

### Lesson 4.2 — Variadic Functions

**Concept slide**

- `...T` last param, behaves like `[]T`; spread with `s...`.

**Tasks**
- Beginner: `Max(nums ...int) int` — 0 if empty.
- Intermediate: `Join(sep string, parts ...string) string`.
- Advanced: `Mean(scores ...float64) float64` with spread.

---

### Lesson 4.3 — Closures

**Concept slide**

- Functions capturing outer scope state; fresh state per call to the generator.

**Tasks**
- Beginner: `MakeCounter() func() int`.
- Intermediate: `MakeAdder(add int) func(int) int`.
- Advanced: `MakeBank(balance float64) func(float64) float64`.

---

### Lesson 4.4 — `defer`

**Concept slide**

- Runs before return; LIFO order; cleanup idiom.

**Tasks**
- Beginner: `ProcessLog() []string` — one defer appends `"end"`.
- Intermediate: `StackLog() []string` — two defers prove LIFO.
- Advanced: `LoopCleanup(n int) []int` — per-loop defer reverses order.

---

## 🧩 Chapter 4 Quiz (5 questions, ~2 min)

1. What does `...` in a parameter list mean? *(variadic)*
2. Predict the output:
   ```go
   defer fmt.Println("A")
   fmt.Println("B")
   ```
   *(B then A)*
3. What is a closure? *(function capturing surrounding-scope variables)*
4. How many values can a Go function return? *(as many as declared — commonly `(value, error)`)*
5. True/False: `defer` statements run in declaration order (FIFO). *(False — LIFO)*