# Chapter 4 Quiz — Functions

Five quick questions (~2 min). Answers are shown on the next screen after the student responds.

---

**Q1. What does `...` in a parameter list mean?**

- A) The loop continues forever
- B) The parameter is variadic — it accepts any number of arguments
- C) The function is private

**Answer:** B — `func sum(nums ...int)` accepts zero or more `int` arguments; inside, `nums` is a `[]int`.

---

**Q2. Predict the output:**
```go
defer fmt.Println("A")
fmt.Println("B")
```

- A) `A B`
- B) `B A`
- C) `A`

**Answer:** B — the `defer` runs right before the function returns, so `"B"` prints first, then `"A"`.

---

**Q3. What is a closure?**

- A) A function that captures variables from the scope it was created in
- B) A function with no return value
- C) A private method

**Answer:** A — closures "close over" surrounding variables and keep them alive after the outer function returns.

---

**Q4. How many values can a Go function return?**

- A) Exactly one
- B) At most two
- C) As many as declared — commonly `(value, error)`

**Answer:** C — Go functions can return any declared number of values; the `(value, error)` pair is the idiomatic pattern.

---

**Q5. True/False: `defer` statements run in the order they were declared (FIFO).**

- A) True
- B) False

**Answer:** B — defers run in **LIFO** order (last declared, first executed), like a stack.