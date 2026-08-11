> A café charges a flat fare. Complete `ClassFare(age int, isStudent bool) int`:
> - children under 12 and seniors 65+ pay `25`
> - students pay `30` (unless they already qualify for the `25` fare)
> - everyone else pays `50`
>
> Use an `if` / `else if` chain with `&&`/`||`, and check the special fares first.
>
> **Expected behavior:**
> ```go
> ClassFare(8, false)   // 25
> ClassFare(70, false)  // 25
> ClassFare(20, true)   // 30
> ClassFare(20, false)  // 50
> ClassFare(10, true)   // 25 (age < 12 qualifies first — special fare wins)
> ```