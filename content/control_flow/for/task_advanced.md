> Complete `IsPrime(n int) bool`. Return `true` if `n` is prime (greater than 1 and divisible only by 1 and itself), and `false` otherwise. For efficiency, only test divisors from `2` up to `i*i <= n` using a `for` loop with an early `return`.
>
> **Expected behavior:**
> ```go
> IsPrime(2)  // true
> IsPrime(7)  // true
> IsPrime(8)  // false
> IsPrime(1)  // false
> ```