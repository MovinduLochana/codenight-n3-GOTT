> Define `WeightedRect` that embeds `Rectangle` and adds a `Weight float64` field. Give it its **own** `Area()` method that returns the promoted `Rectangle` area multiplied by `Weight` (i.e. `nr.Rectangle.Area() * nr.Weight`). This shadows the promoted method on the outer type.
>
> **Expected behavior:**
> ```go
> wr := WeightedRect{Rectangle: Rectangle{Width: 3, Height: 4}, Weight: 2}
> wr.Area()   // 12 * 2 = 24 (WeightedRect's own method)
> // promoted access still works:
> wr.Rectangle.Area() // 12
> ```