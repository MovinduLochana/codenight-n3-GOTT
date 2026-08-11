> Complete `Invert(m map[string]int) map[int][]string`. It should return a map that swaps keys and values: each original value becomes a key, and its slice holds all the original keys that mapped to it. Build the `[]string` with `make`/`append`.
>
> **Expected behavior (order of each value slice not guaranteed — hidden test deep-compares):**
> ```go
> Invert(map[string]int{"a": 1, "b": 1, "c": 2})
> // map[1:[a b] 2:[c]]  (order of a/b varies)
> ```