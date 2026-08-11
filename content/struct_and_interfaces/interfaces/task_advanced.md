> Complete `Biggest(shapes []Shape) float64`. It should return the **largest** area among all shapes in the slice. Use a `range` loop over the `Shape` interface values. Assume the slice is non-empty.
>
> **Expected behavior:**
> ```go
> shapes := []Shape{
>     Rectangle{Width: 2, Height: 3}, // area 6
>     Circle{Radius: 2},              // area ≈ 12.56636
> }
> Biggest(shapes) // ≈ 12.56636
> ```