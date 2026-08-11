> Define `Square` that embeds `Rectangle` and adds a `Side float64` field. Its **own** `Area()` method should ignore the promoted fields and return `Side * Side`. `Rectangle` already has `Area()`; also define `func TotalArea(shapes []Shape) float64` summing `Area()` over a `Shape` interface. Because `Square` gets a promoted `Rectangle.Area()` method automatically, a `Square` value satisfies `Shape` too — but its own shadowing method wins.
>
> **Expected behavior:**
> ```go
> shapes := []Shape{
>     Square{Rectangle: Rectangle{Width: 3, Height: 4}, Side: 5},
> }
> TotalArea(shapes) // 25 (Square.Area — its own method, not the promoted one)
> ```