> Complete `ProcessLog() (log []string)`. Append `"start"` and `"middle"` normally, then use exactly **one** `defer` to append `"end"` **after** them, using a **named return value** and a bare `return`.
>
> > **Key point:** the deferred append must go through the *named* return value (`log`). If you used `return log` with an anonymous return, the return value would be snapshotted *before* the deferred call runs and `"end"` would be lost.
>
> **Expected behavior:**
> ```go
> ProcessLog() // ["start", "middle", "end"]
> ```