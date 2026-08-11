> Define an interface `Pricer` with one method `Price() float64`. Give `Product` a `Price() float64` method that returns `p.BasePrice`. Then complete `TotalPrice(items []Pricer) float64` which sums every item's price. This is "many products, one interface" — the same pattern as `Shape`/`Area()`. (Note: the stored amount is a field named `BasePrice`, because a struct can't have both a field and a method named `Price`.)
>
> **Expected behavior:**
> ```go
> items := []Pricer{
>     Product{Name: "Coffee", BasePrice: 4.50},
>     Product{Name: "Cake", BasePrice: 6.00},
> }
> TotalPrice(items) // 10.5
> ```