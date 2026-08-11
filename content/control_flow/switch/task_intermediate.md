> Complete `DaysInMonth(m int) int` using a `switch` statement with grouped cases. Return the number of days for the given month number (`1`=January … `12`=December):
> - February → `28`
> - April, June, September, November → `30`
> - all others → `31`
> - an invalid month → `0`
>
> **Expected behavior:**
> ```go
> DaysInMonth(2)  // 28
> DaysInMonth(4)  // 30
> DaysInMonth(1)  // 31
> DaysInMonth(13) // 0
> ```