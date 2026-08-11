> Complete `ReadAge(s string) (int, error)`. Use `strconv.Atoi` to parse `s`. On a parse error, **wrap** it with `fmt.Errorf("invalid age: %w", err)` — this preserves the original cause. On success return the number and `nil`.
>
> Then, in `main`, use `errors.Is(err, strconv.ErrSyntax)` to detect a syntax (parse) failure through the wrapping. `%w` is what makes this detection possible.
>
> **Expected behavior:**
> ```go
> ReadAge("25") // 25, nil
> ReadAge("xx") // 0, error — errors.Is(err, strconv.ErrSyntax) == true
> ```