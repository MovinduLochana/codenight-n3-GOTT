# Chapter 3 — Collections: Arrays, Slices, Maps (~30 min)

Format per lesson: **one concept slide → three tasks** (Beginner → Intermediate → Advanced).

| # | Lesson | Time |
| --- | -------- | ------ |
| 3.1 | Arrays vs. Slices | 7 min |
| 3.2 | Slicing, `append`, and the Shared Backing Array | 8 min |
| 3.3 | Maps | 7 min |
| 3.4 | `range` over Slices & Maps | 6 min |
| — | Chapter 3 Quiz (5 Qs) | 2 min |
| | **Total** | **~30 min** |

---

### Lesson 3.1 — Arrays vs. Slices

**Concept slide**

- Fixed-size `[N]T` arrays vs resizable `[]T` slices.
- `len`/`cap`; `append` grows slices.
- One idea only: use slices, almost never arrays.

**Tasks**
- Beginner: `AppendThree(s []int) []int` — append `1, 2, 3`.
- Intermediate: `RangeOf(n int) []int` — build `[0..n)` with `make` + loop.
- Advanced: `Tail(s []int, n int) []int` — copy last `n` elements.

---

### Lesson 3.2 — Slicing `[low:high]`, `append`, backing arrays

**Concept slide**

- `s[low:high]` sub-slices; shared backing array; `copy()`.
- One idea only: slicing views share memory.

**Tasks**
- Beginner: `Cut(s []int, lo, hi int) []int` — plain `s[lo:hi]` view.
- Intermediate: `RemoveAt(s []int, i int) []int` — `append(s[:i], s[i+1:]...)` (source, backing-array gotcha).
- Advanced: `InsertAt(s []int, i, v int) []int` — order-preserving insert.

---

### Lesson 3.3 — Maps

**Concept slide**

- `make(map[K]V)`, set/get, `delete`, comma-ok, nil-map write panic.

**Tasks**
- Beginner: `WordCount(words []string) map[string]int`.
- Intermediate: `MergeCounts(a, b map[string]int) map[string]int` — comma-ok sums.
- Advanced: `Invert(m map[string]int) map[int][]string`.

---

### Lesson 3.4 — `range`

**Concept slide**

- `for i, v := range`, discard index with `_`, map order is randomized.

**Tasks**
- Beginner: `Sum(nums []int) int` — `for _, v := range`.
- Intermediate: `Contains(nums []int, target int) bool`.
- Advanced: `IndexAll(nums []int, target int) []int` — uses range index.

---

## 🧩 Chapter 3 Quiz (5 questions, ~2 min)

1. What's the zero value of an uninitialized `map[string]int`? *(nil)*
2. What does the "comma ok" idiom check? *(whether the key exists)*
3. Predict the output:
   ```go
   s := []int{1, 2, 3}
   s = append(s, 4)
   fmt.Println(len(s))
   ```
   *(4)*
4. True/False: iterating a Go map always visits keys in insertion order. *(False)*
5. What function grows a slice? *(`append`)*