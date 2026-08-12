> 1. Create a `BankAccount` struct with two fields: `Owner` (string) and `Balance` (float64).
> 2. Implement a method `Deposit(amount float64)` with a **pointer receiver** that adds the `amount` to the account's balance.
> 3. Implement a method `GetDetails() string` with a **value receiver** that returns a formatted string like: `"Account: [Owner], Balance: $[Balance]"`. Format the balance to 2 decimal places.
>
> **Expected Output:**
>
> ```
> Account: Alice, Balance: $150.50
> ```
