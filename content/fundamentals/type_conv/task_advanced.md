> Complete the function `ToPixels(inches, dpi float64) int`. It should convert a measurement like `1.5` inches at `300` DPI into an integer pixel count by multiplying **in float64 first**, then converting the result to `int` with `int(...)`. This demonstrates converting float → int (which **truncates** the decimal part).
>
> **Expected behavior:**
> ```go
> ToPixels(1.5, 300) // 450
> ToPixels(0.5, 72)  // 36
> ```