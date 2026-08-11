> Add a method `Scale(factor float64)` on `*Rectangle` (pointer receiver) that multiplies both `Width` and `Height` by `factor`, mutating the original struct.
>
> **Expected behavior:**
> ```go
> r := Rectangle{Width: 3, Height: 4}
> r.Scale(2)
> // r.Width == 6, r.Height == 8
> ```