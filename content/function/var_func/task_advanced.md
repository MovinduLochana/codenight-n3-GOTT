> Complete `Mean(scores ...float64) float64`. It should return the arithmetic average of all scores. If no scores are given, return `0`. In `main`, also call `Mean` by **spreading a slice** with `values...` to prove the two forms are equivalent.
>
> **Expected behavior:**
> ```go
> Mean(90, 86, 88)    // 88
> Mean(10, 20)        // 15
> Mean()              // 0
> values := []float64{2, 4} ; Mean(values...) // 3
> ```