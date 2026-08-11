# Chapter 1 Quiz — Go Fundamentals

Five quick questions (~2 min). Answers are shown on the next screen after the student responds.

---

**Q1. Which two forms of variable declaration did you learn first?**

- A) `var` and `:=`
- B) `let` and `const`
- C) `define` and `:`

**Answer:** A — `var` is the explicit form; `:=` is the short declaration that infers the type.

---

**Q2. What's the zero value of a `string`?**

- A) `nil`
- B) `"null"`
- C) `""`

**Answer:** C — an uninitialized string's zero value is the empty string `""`. `nil` is the zero value for maps, slices, pointers, and interfaces.

---

**Q3. What keyword declares a compile-time constant?**

- A) `let`
- B) `const`
- C) `static`

**Answer:** B — `const = <value>` declares a value fixed at compile time; `iota` makes sequential constants easy.

---

**Q4. True/False: Go auto-converts `int` to `float64` when needed.**

- A) True
- B) False

**Answer:** B — Go never converts types automatically. Every conversion is explicit, e.g. `float64(i)`.

---

**Q5. Which `fmt` verb prints the type of a value?**

- A) `%v`
- B) `%T`
- C) `%t`

**Answer:** B — `%T` prints the type, e.g. `fmt.Printf("%T", age)` → `int`.