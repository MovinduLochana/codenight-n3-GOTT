> Implement a function `CountLines(filename string) (int, error)` that reads a text file and returns the total count of lines.
> - The function should open the file, scan it line-by-line using `bufio.NewScanner`, increment a counter, and return the final count.
> - Make sure to close the file properly (using `defer`).
> - If the file does not exist or fails to open, return `0` and the corresponding error.
> - In `main()`, print the line count of `"sample.txt"`.
>
> **Expected Output:**
>
> ```
> Line count: 3
> ```
