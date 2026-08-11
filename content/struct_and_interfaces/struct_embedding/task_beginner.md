> Define `NamedRectangle` with an embedded `Rectangle` and an additional field `Name string`. Then, without writing a new `Area()` method, call the promoted `Area()` method directly on a `NamedRectangle` value.
>
> **Expected behavior:**
> ```go
> nr := NamedRectangle{Rectangle: Rectangle{Width: 3, Height: 4}, Name: "MyRect"}
> nr.Area() // 12 (promoted, not redefined)
> ```