# Chapter 5 — Structs & Interfaces (~35 min)

Format per lesson: **one concept slide → three tasks** (Beginner → Intermediate → Advanced).

> **Flow note:** Lesson 5.2 (Pointers) was added to this chapter so that pointer receivers in 5.3 build on a solid foundation — `&`/`*` are used everywhere in Go but were never introduced in Chapters 1–4.

| # | Lesson | Time |
| --- | -------- | ------ |
| 5.1 | Structs | 5 min |
| 5.2 | Pointers (`&` / `*`) | 6 min |
| 5.3 | Methods & Pointer Receivers | 6 min |
| 5.4 | Interfaces | 6 min |
| 5.5 | Struct Embedding | 5 min |
| 5.6 | Capstone — Café Menu | 5 min |
| — | Chapter 5 Quiz (5 Qs) | 2 min |
| | **Total** | **~35 min** |

---

### Lesson 5.1 — Structs

**Concept slide**

- `type X struct { A T; B T }`; literals by name or position; plain data containers.

**Tasks**
- Beginner: `Rectangle{Width, Height}` + `Area(r) float64`.
- Intermediate: `Box{Length, Width, Height}` + `Volume(b) float64`.
- Advanced: `Product{Name, Price}` + slice field `Cart{items}` + `AddItem`.

---

### Lesson 5.2 — Pointers

**Concept slide**

- `&x` address-of; `*p` deref; mutating through pointers; why Go uses them (avoid copies, shared state, pointer receivers).

**Tasks**
- Beginner: `Set10(ptr *int)` — `*ptr = 10`.
- Intermediate: `Double(ptr *int) int` — `*ptr *= 2`, return new value.
- Advanced: `Swap(a, b *int)` — exchange pointed-to values.

---

### Lesson 5.3 — Methods & Pointer Receivers

**Concept slide**

- Value receiver = copy; pointer receiver = mutate original; auto `&r` when calling a pointer method.

**Tasks**
- Beginner: `Scale(factor float64)` on `*Rectangle`.
- Intermediate: named type `Money float64` with a value-receiver method.
- Advanced: `Stack` with pointer `Push`/`Pop` managing a slice field.

---

### Lesson 5.4 — Interfaces

**Concept slide**

- Implicit satisfaction; behavior-based polymorphism; no `implements`.

**Tasks**
- Beginner: `Circle` + `TotalArea(shapes []Shape)`.
- Intermediate: second interface `Perimeterer` + `TotalPerimeter`.
- Advanced: `Biggest(shapes []Shape) float64`.

---

### Lesson 5.5 — Struct Embedding

**Concept slide**

- Composition via embedding; promoted fields/methods; "has-a", not "is-a".

**Tasks**
- Beginner: `NamedRectangle` embedding `Rectangle`.
- Intermediate: `WeightedRect` overriding logic using promoted field.
- Advanced: embedded type satisfies interface in `[]Shape`.

---

### Lesson 5.6 — Capstone: Café Menu

**Concept slide (practice only, no new concept)**

- Combine structs, methods, interfaces, embedding into a tiny menu app.

**Tasks**
- Beginner: `Product{Name, Price}.Label() string`.
- Intermediate: `Pricer` interface + `TotalPrice([]Pricer) float64`.
- Advanced: embedded `SpecialItem` overriding `Price()` with discount.

---

## 🧩 Chapter 5 Quiz (5 questions, ~2 min)

1. Does Go have classes? *(No — structs + methods)*
2. What does `*` do in `*p`? *(dereferences — accesses the pointed-to value)*
3. When must you use a pointer receiver? *(mutating the receiver, or big structs)*
4. How does a type satisfy an interface? *(implicitly, by implementing its methods)*
5. What is struct embedding for? *(composition — promoting fields/methods)*