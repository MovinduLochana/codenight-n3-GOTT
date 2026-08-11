> Using an `iota`-based `const` block, define four constants `Red`, `Green`, `Blue`, `Alpha` **in that order**, but skip the first `iota` value (`0`) by assigning it to the blank identifier `_`:
>
> ```go
> const (
>     _ = iota
>     Red
>     Green
>     Blue
>     Alpha
> )
> ```
>
> Complete `BlueValue()` so it returns the numeric value of `Blue` (the starter already has the `_` skip in place for you).
>
> **Expected behavior:**
> ```go
> BlueValue() // returns 3
> ```