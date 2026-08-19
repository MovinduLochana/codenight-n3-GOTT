> Fill in the `Greet` function so it reads the user's name with `fmt.Scanln` and returns the greeting:
>
> ```
> "Hello, <name>!"
> ```
>
> **Example:**
> Input: `Gopher` → returns `"Hello, Gopher!"`
>
> **Hints:**
> - Declare a `var name string`.
> - Use `fmt.Scanln(&name)` to read one line of input.
> - Return `"Hello, " + name + "!"`.