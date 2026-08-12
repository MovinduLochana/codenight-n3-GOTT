> Implement a worker pool pattern using Goroutines, Channels, and a WaitGroup.
> 1. Implement a function `worker(id int, jobs <-chan int, results chan<- int)`:
>    - Reads jobs from the `jobs` channel.
>    - Multiplies each job integer by `2`.
>    - Sends the result to the `results` channel.
> 2. In `main()`:
>    - Initialize two buffered channels `jobs` and `results` of size `5`.
>    - Spawn `3` worker goroutines using `sync.WaitGroup` to track their execution.
>    - Send `5` jobs (integers `1` to `5`) to the `jobs` channel and then close it.
>    - Ensure that the `results` channel is closed exactly when all workers have finished processing.
>    - Read the results from the `results` channel, sort them, and print them.
>
> **Expected Output:**
>
> ```
> Results: [2 4 6 8 10]
> ```
