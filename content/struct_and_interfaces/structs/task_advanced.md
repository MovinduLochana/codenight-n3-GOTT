> Define a struct `Product` with fields `Name string` and `Price float64`, and a struct `Cart` with a slice field `Items []Product`. Then complete `AddItem(c Cart, p Product) Cart` — it should return a **new** `Cart` with `p` appended to `c.Items` (copy the slice header; no pointers needed yet).
>
> **Expected behavior:**
> ```go
> c := Cart{}
> c = AddItem(c, Product{Name: "Coffee", Price: 4.50})
> c = AddItem(c, Product{Name: "Tea", Price: 3.00})
> len(c.Items) // 2
> c.Items[0].Name // "Coffee"
> ```