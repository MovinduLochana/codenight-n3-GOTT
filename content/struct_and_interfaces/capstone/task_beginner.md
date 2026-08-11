> Define a struct `Product` with fields `Name string` and `Price float64`. Give it a value-receiver method `Label() string` that returns `"<Name>: $<Price to 2 decimals>"` using `fmt.Sprintf`.
>
> **Expected behavior:**
> ```go
> Product{Name: "Coffee", Price: 4.50}.Label() // "Coffee: $4.50"
> Product{Name: "Tea", Price: 3.00}.Label()    // "Tea: $3.00"
> ```