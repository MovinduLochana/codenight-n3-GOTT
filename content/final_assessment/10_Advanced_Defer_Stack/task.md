> 1. Implement `Push` and `Pop` methods on the `Stack` struct (using a pointer receiver).
>    - `Push(val string)` appends a string to the elements slice.
>    - `Pop() (string, bool)` removes the top string from the stack and returns it along with `true`. If the stack is empty, return `""` and `false`.
> 2. In `main()`, register a `defer` block that will continually pop and print all elements from the stack in LIFO (Last In First Out) order when `main()` completes.
>
> **Expected Output:**
>
> ```
> Beginning transaction...
> Transaction steps registered.
> Rolling back or auditing transaction steps:
> - COMMIT
> - WRITE_TEMP
> - START
> ```
