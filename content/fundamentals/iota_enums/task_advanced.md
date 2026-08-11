> `iota` isn't just for counting to 1, 2, 3 — you can combine it with operators. Using a `const` block, define three flag constants as **powers of two** using `1 << iota`:
>
> ```go
> const (
>     Read = 1 << iota // 1
>     Write            // 2
>     Execute          // 4
> )
> ```
>
> Then complete the function `AllPermissions() int` so it returns `Read + Write + Execute`.
>
> **Expected behavior:**
> ```go
> AllPermissions() // returns 7
> ```