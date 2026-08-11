> Complete the function `IsLeapYear(year int) bool`. A year is a leap year if it is divisible by 4 — EXCEPT century years (divisible by 100), which are leap years only if also divisible by 400. Combine `%`, `==`, `&&`, and `||`.
>
> **Expected behavior:**
> ```go
> IsLeapYear(2024) // true  (divisible by 4)
> IsLeapYear(1900) // false (century, not divisible by 400)
> IsLeapYear(2000) // true  (century, divisible by 400)
> ```