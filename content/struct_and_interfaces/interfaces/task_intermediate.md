> Define an interface `Perimeterer` with a method `Perimeter() float64`. Give `Rectangle` and `Circle` `Perimeter() float64` methods:
> - `Rectangle`: `2 * (Width + Height)`
> - `Circle`: `2 * 3.14159 * Radius`
>
> Then complete `TotalPerimeter(shapes []Perimeterer) float64`, which sums every shape's perimeter. Note that both types implicitly satisfy *both* `Shape` and `Perimeterer`.
>
> **Expected behavior:**
> ```go
> shapes := []Perimeterer{Rectangle{Width: 2, Height: 3}, Circle{Radius: 1}}
> TotalPerimeter(shapes) // 10 + 6.28318 ≈ 16.28318
> ```