> Complete `Season(m int) string` using an **expression-less `switch`** with grouped multi-value cases. Given a month number (`1`=January … `12`=December), return:
> - `"Winter"` for `12`, `1`, `2`
> - `"Spring"` for `3`, `4`, `5`
> - `"Summer"` for `6`, `7`, `8`
> - `"Fall"` for `9`, `10`, `11`
> - `"invalid"` for anything else
>
> **Expected behavior:**
> ```go
> Season(12) // "Winter"
> Season(4)  // "Spring"
> Season(8)  // "Summer"
> Season(10) // "Fall"
> Season(0)  // "invalid"
> ```