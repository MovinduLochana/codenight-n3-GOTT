> Define a struct named `Product` with fields:
> - `ID` (int)
> - `Name` (string)
> - `Price` (float64)
> - `InStock` (bool)
>
> Add struct tags to ensure the JSON keys are lowercase and snake_case: `id`, `name`, `price`, `in_stock`.
>
> Implement a function `SerializeProduct(p Product) (string, error)` that marshals a product struct to a JSON string.
> In `main()`, serialize a product with `ID: 101`, `Name: "Wireless Mouse"`, `Price: 29.99`, and `InStock: true`, then print the JSON string.
>
> **Expected Output:**
>
> ```
> {"id":101,"name":"Wireless Mouse","price":29.99,"in_stock":true}
> ```
