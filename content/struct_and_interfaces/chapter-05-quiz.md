# Chapter 5 Quiz — Structs & Interfaces

Five quick questions (~2 min). Answers are shown on the next screen after the student responds.

---

**Q1. Does Go have classes?**

- A) Yes, with inheritance
- B) No — structs + methods
- C) Yes, but only virtual

**Answer:** B — Go has no classes and no class inheritance; you model data with **structs** and attach behavior with **methods**.

---

**Q2. What does `*` do in `*p`?**

- A) Takes the address of `p`
- B) Dereferences — accesses the pointed-to value
- C) Multiplies `p` by 2

**Answer:** B — `*p` dereferences the pointer; `&x` (not `*p`) takes an address.

---

**Q3. When must you use a pointer receiver?**

- A) When mutating the receiver's fields (or the struct is large)
- B) Always
- C) Only for interfaces

**Answer:** A — use a pointer receiver when a method must mutate the receiver, or to avoid copying a large struct; a value receiver works on a copy.

---

**Q4. How does a type satisfy an interface?**

- A) By declaring `implements`
- B) Implicitly, by implementing the interface's methods
- C) With a field of the interface type

**Answer:** B — Go interfaces are satisfied implicitly; if the method set matches, it just works. No `implements` keyword.

---

**Q5. What is struct embedding for?**

- A) Inheritance ("is-a")
- B) Composition — promoting fields and methods into the outer struct
- C) Encapsulation only

**Answer:** B — embedding promotes an inner struct's fields/methods into the outer one ("has-a"). It's composition, not inheritance.