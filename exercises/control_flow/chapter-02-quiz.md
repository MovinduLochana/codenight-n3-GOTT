# Chapter 2 Quiz — Control Flow

Five quick questions (~2 min). Answers are shown on the next screen after the student responds.

---

**Q1. Does a Go `switch` case fall through by default?**

- A) Yes
- B) No
- C) Only with the `case` keyword

**Answer:** B — Go `switch` cases do **not** fall through; only the matching case runs (use explicit `fallthrough` to override).

---

**Q2. Predict the output:**
```go
for i := 0; i < 3; i++ {
    if i == 1 {
        continue
    }
    fmt.Print(i)
}
```

- A) `012`
- B) `02`
- C) `0`

**Answer:** B — `continue` skips the `i == 1` iteration, so only `0` and `2` print.

---

**Q3. What are the three clauses in a classic `for` loop called?**

- A) start; stop; step
- B) init; condition; post
- C) setup; check; update

**Answer:** B — `for init; condition; post { ... }`.

---

**Q4. Is `while` a keyword in Go?**

- A) Yes, Go has `while` and `for`
- B) No — `for` covers it
- C) Yes, `while` is an alias for `for`

**Answer:** B — Go has exactly one loop keyword: `for`. A condition-only `for cond { }` works as a `while` loop.

---

**Q5. What keyword exits a loop early?**

- A) `break`
- B) `exit`
- C) `stop`

**Answer:** A — `break` exits a loop entirely; `continue` skips to the next iteration.