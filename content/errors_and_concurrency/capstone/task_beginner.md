> Complete `Validate(item string, qty int) error`. Return an error (using `errors.New` or `fmt.Errorf`) when `qty < 1` — an error **is** the return value here. Return `nil` when the quantity is valid. This is the "errors as values" pattern from Lesson 6.1.
>
> **Expected behavior:**
> ```go
> Validate("Coffee", 2) // nil
> Validate("Tea", 0)    // error
> Validate("Tea", -3)   // error
> ```