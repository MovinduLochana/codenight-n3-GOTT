> Define a type `Money float64` with a **value receiver** method `Format()` that returns the amount as a `float64` (just returns `float64(m)`). Then define `func AddCurrency(a, b Money) Money` that adds two `Money` values. The point: `Money` is a named type with methods, not a struct.
>
> **Expected behavior:**
> ```go
> m := Money(12.5)
> m.Format()        // 12.5
> AddCurrency(Money(1.25), Money(2.50)) // Money(3.75)
> ```