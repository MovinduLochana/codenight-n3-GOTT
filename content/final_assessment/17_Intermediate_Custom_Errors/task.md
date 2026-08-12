> 1. Create a custom error struct type named `ErrInvalidAge` that implements the `error` interface. It should store the field:
>    - `Age` (int)
>    - The `Error()` method should return a string like `"age [Age] is invalid"`.
> 2. Implement a function `VerifyAge(age int) error`.
>    - If the `age` is negative (`< 0`) or greater than `150`, return a pointer to `ErrInvalidAge` configured with the invalid age.
>    - Otherwise, return `nil`.
> 3. In `main()`, call `VerifyAge(-5)`. If an error occurs, check if it's of type `*ErrInvalidAge` using `errors.As` and print:
>    `Error: age -5 is invalid`
>
> **Expected Output:**
>
> ```
> Error: age -5 is invalid
> ```
