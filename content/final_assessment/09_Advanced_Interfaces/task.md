> 1. Define an interface `Item` with a method `Price() float64`.
> 2. Create a struct `BaseItem` with fields `Name` (string) and `BasePrice` (float64) that implements the `Item` interface.
> 3. Create a struct `TaxableItem` that embeds `BaseItem` and overrides the `Price()` method to add an additional 15% tax to the `BasePrice`.
> 4. Implement a function `CalculateTotal(items []Item) float64` that calculates the total price of all items in the slice.
>
> **Expected Output:**
>
> ```
> Total Cart Price: $1180.00
> ```
