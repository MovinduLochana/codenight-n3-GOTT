> Refactor `ParsePositive(s string) (int, error)` to use a **sentinel error**. Declare a package-level `var ErrNotPositive = errors.New("number must be positive")`. Parse `s` with `strconv.Atoi` (return the raw parse error unwrapped), and if the number is **zero or negative**, return `0, ErrNotPositive`. Otherwise return the number and `nil`. In `main`, detect the sentinel with `errors.Is(err, ErrNotPositive)`.
>
> **Expected behavior:**
> ```go
> ParsePositive("7")                       // 7, nil
> ParsePositive("-5")                      // 0, ErrNotPositive (errors.Is == true)
> ParsePositive("0")                       // 0, ErrNotPositive (errors.Is == true)
> ParsePositive("abc")                     // 0, parse error (errors.Is == false)
> ```