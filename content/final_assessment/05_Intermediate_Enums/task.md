> 1. Define a custom type named `UserRole` with `int` as its underlying type.
> 2. Declare a group of constants using `iota` representing:
>    - `Guest` (value 0)
>    - `Member` (value 1)
>    - `Admin` (value 2)
>    - `Owner` (value 3)
> 3. Implement a function `GetPermission(role UserRole) string` that returns permissions as follows using a `switch` statement:
>    - `Guest` -> `"Read Only"`
>    - `Member` -> `"Read & Write"`
>    - `Admin` -> `"Read, Write & Moderate"`
>    - `Owner` -> `"Full Access"`
>    - Any other value -> `"Unknown"`
>
> **Expected Output:**
>
> ```
> Guest: Read Only
> Member: Read & Write
> Admin: Read, Write & Moderate
> Owner: Full Access
> ```
