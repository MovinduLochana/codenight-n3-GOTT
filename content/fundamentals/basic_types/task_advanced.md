> Combine the four basic types in one small program. Declare:
> - `name := "Gopher"` — string
> - `visits := 7` — int
> - `renewed := visits > 3` — a `bool` produced by a comparison
> - `rating := 4.75` — float64
>
> Print one line: `"Welcome back " + name`, then `visits`, then `renewed`, then `rating` — all space-separated with `fmt.Println`.
>
> **Expected Output:**
>
> ```
> Welcome back Gopher 7 true 4.75
> ```