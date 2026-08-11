> Define a `Circle` struct with field `Radius float64` and give it an `Area() float64` method (use `3.14159` for π). `Rectangle` and its `Area()` method are already provided. Then complete `TotalArea(shapes []Shape) float64`, which sums the area of every shape in the slice.
>
> **Expected behavior:**
> ```go
> shapes := []Shape{Rectangle{Width: 2, Height: 3}, Circle{Radius: 1}}
> TotalArea(shapes) // 6 + 3.14159 ≈ 9.14159
> ```