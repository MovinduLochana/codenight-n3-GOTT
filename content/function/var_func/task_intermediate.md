> Complete `Join(sep string, parts ...string) string`. It should concatenate all `parts` with `sep` between them, using a `range` loop and the `+` operator. It should never place a separator before the first or after the last part.
>
> **Expected behavior:**
> ```go
> Join(", ", "go", "is", "fun") // "go, is, fun"
> Join("|", "a")                // "a"
> Join("-")                     // ""
> ```