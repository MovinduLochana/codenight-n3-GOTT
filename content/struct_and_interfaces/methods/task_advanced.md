> Define a type `Stack` with a field `Items []int`. Add **pointer-receiver** methods that mutate the stack:
> - `Push(v int)` — appends `v` to the stack
> - `Pop() int` — removes and returns the **last** item
>
> A **pointer receiver** is required because both methods mutate the receiver (including `Pop`, which must shrink the slice and return the removed value). You can assume `Pop` is only called on a non-empty stack.
>
> **Expected behavior:**
> ```go
> s := Stack{}
> s.Push(1); s.Push(2); s.Push(3)
> s.Pop() // 3
> s.Pop() // 2
> s.Items // [1]
> ```