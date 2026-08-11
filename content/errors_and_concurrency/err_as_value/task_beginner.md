> Complete `ParsePositive(s string) (int, error)`. Use `strconv.Atoi` to parse `s`. Return an error if parsing fails, or if the parsed number is negative. Otherwise return the number and `nil`.
>
> **Expected behavior:**
> ```go
> ParsePositive("42")  // 42, nil
> ParsePositive("-5")  // 0, error
> ParsePositive("abc") // 0, error
> ```