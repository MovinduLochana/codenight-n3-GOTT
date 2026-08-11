# Chapter 2 — Control Flow (~20 min)

Format per lesson: **one concept slide → three tasks** (Beginner → Intermediate → Advanced). No lesson introduces more than one new idea.

| # | Lesson | Time |
| --- | -------- | ------ |
| 2.1 | `if` / `else` and Init Statements | 7 min |
| 2.2 | The One True Loop: `for` | 7 min |
| 2.3 | `switch` | 6 min |
| — | Chapter 2 Quiz (5 Qs) | 2 min |
| | **Total** | **~22 min** |

---

### Lesson 2.1 — `if` / `else` and Init Statements

**Concept slide**

- `if <condition> { ... } else if { ... } else { ... }` — no parens needed, braces mandatory.
- `if <init>; <condition> { ... }` — init var scoped to the block.
- One idea only: branching on conditions.

**Tasks**
- Beginner: `Grade(score int) string` → `"A"`/`"B"`/`"C"`/`"F"`.
- Intermediate: `Sign(n int) string` → `"positive"`/`"negative"`/`"zero"`.
- Advanced: `ClassFare(age int, isStudent bool) int` — nested `if`/`else` with `&&`/`||`.

---

### Lesson 2.2 — The One True Loop: `for`

**Concept slide**

- `for init; cond; post {}`, `for cond {}` (while), `for {}` (infinite).
- `break` exits, `continue` skips.
- One idea only: Go has exactly one loop keyword.

**Tasks**
- Beginner: `SumEvens(n int) int` — sum evens `1..n`.
- Intermediate: `Factorial(n int) int` — `n!`, `Factorial(0) == 1`.
- Advanced: `IsPrime(n int) bool` — trial division loop up to `i*i <= n`.

---

### Lesson 2.3 — `switch`

**Concept slide**

- `switch expr { case ... default: }`, expression-less `switch {}` chains.
- No fallthrough by default; `fallthrough` is rare.

**Tasks**
- Beginner: `DayType(day string) string` — `"Weekend"` vs `"Weekday"`.
- Intermediate: `DaysInMonth(m int) int` — grouped cases Feb/30-day/31-day.
- Advanced: `Season(m int) string` — expression-less switch, multi-value cases (`case 12, 1, 2`).

---

## 🧩 Chapter 2 Quiz (5 questions, ~2 min)

1. Does a Go `switch` case fall through by default? *(No)*
2. Predict the output:
   ```go
   for i := 0; i < 3; i++ {
       if i == 1 { continue }
       fmt.Print(i)
   }
   ```
   *(02)*
3. What are the three clauses in a classic `for` loop called? *(init; condition; post)*
4. Is `while` a keyword in Go? *(No — `for` covers it)*
5. What keyword exits a loop early? *(`break`)*