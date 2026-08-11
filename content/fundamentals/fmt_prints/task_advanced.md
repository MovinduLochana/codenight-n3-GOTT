> Use `fmt.Printf` to format a single line with **zero-padding and float rounding**. Given the variables `score := 42` and `pi := 3.14159`:
> - print `score` with the `%04d` verb (min width 4, zero-padded → `0042`)
> - print `pi` with the `%.2f` verb (rounded to 2 decimals → `3.14`)
> separated by a single space, ending with a newline.
>
> **Expected Output:**
>
> ```
> 0042 3.14
> ```