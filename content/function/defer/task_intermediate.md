> Complete `StackLog() (log []string)`. Schedule **two** deferred appends on the named return: defer `"B"` **first**, then defer `"A"`. Because defers run **LIFO**, `"A"` (deferred last) must be appended before `"B"`. Append `"start"` normally first. Expected `["start", "A", "B"]`.
>
> **Expected behavior:**
> ```go
> StackLog() // ["start", "A", "B"]
> ```