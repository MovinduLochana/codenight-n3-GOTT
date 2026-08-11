> Complete `MergeCounts(a, b map[string]int) map[string]int`. It should return a new map that combines both input maps by summing the count for every key present in either. Use the **comma-ok idiom** to read from `b` and determine whether a key exists.
>
> **Expected behavior (map order not guaranteed — hidden test compares map equality):**
> ```go
> MergeCounts(
>     map[string]int{"go": 2, "is": 1},
>     map[string]int{"go": 1, "fun": 3},
> ) // map[fun:3 go:3 is:1]
> ```