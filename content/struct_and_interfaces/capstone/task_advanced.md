> Define `Deal` that **embeds** `Product` and adds a `Discount float64`. Give `Deal` its **own** `Price() float64` method that returns the discounted price: `d.Product.BasePrice * (1 - d.Discount)` (so a `0.20` discount means 20% off). Because `Deal` has its own `Price()` method, it satisfies `Pricer` polymorphically — the same `TotalPrice([]Pricer)` from the previous task keeps working.
>
> **Expected behavior:**
> ```go
> items := []Pricer{
>     Product{Name: "Coffee", BasePrice: 4.00},
>     Deal{Product: Product{Name: "Cake", BasePrice: 10.00}, Discount: 0.20},
> }
> TotalPrice(items) // 4.00 + 8.00 = 12.0
> ```