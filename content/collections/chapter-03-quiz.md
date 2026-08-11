# Chapter 3 Quiz — Collections: Arrays, Slices, Maps

Five quick questions (~2 min). Answers are shown on the next screen after the student responds.

---

**Q1. What's the zero value of an uninitialized `map[string]int`?**

- A) An empty map
- B) `nil`
- C) A map that panics on read

**Answer:** B — `var m map[string]int` is `nil` until `make`'d. Writing to a nil map panics; reading safely returns the zero value.

---

**Q2. What does the "comma ok" idiom check?**

- A) Whether the map has any elements
- B) Whether a key exists in the map
- C) Whether the value is the zero value

**Answer:** B — `value, ok := m[k]` — `ok` is `true` only when the key exists.

---

**Q3. Predict the output:**
```go
s := []int{1, 2, 3}
s = append(s, 4)
fmt.Println(len(s))
```

- A) `3`
- B) `4`
- C) `5`

**Answer:** B — `append` grows the slice; it now holds 4 elements.

---

**Q4. True/False: iterating a Go map always visits keys in insertion order.**

- A) True
- B) False

**Answer:** B — map iteration order is randomized by design; never rely on it.

---

**Q5. What function grows a slice?**

- A) `grow`
- B) `extend`
- C) `append`

**Answer:** C — `append(s, v...)` returns a larger slice, reallocating the backing array when it runs out of capacity.