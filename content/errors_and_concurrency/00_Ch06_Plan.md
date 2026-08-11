# Chapter 6 — Errors & a Taste of Concurrency (~25 min)

Format per lesson: **one concept slide → three tasks** (Beginner → Intermediate → Advanced).

| # | Lesson | Time |
| --- | -------- | ------ |
| 6.1 | Errors as Values | 7 min |
| 6.2 | Goroutines & WaitGroup | 7 min |
| 6.3 | Channels & Shared State | 7 min |
| 6.4 | Capstone — Concurrent Checkout | 4 min |
| — | Chapter 6 Quiz (5 Qs) | 2 min |
| | **Total** | **~27 min** |

---

### Lesson 6.1 — Errors as Values

**Concept slide**

- `errors.New`, `fmt.Errorf` + `%w` wrapping, `errors.Is`; callers check immediately.

**Tasks**
- Beginner: `ParsePositive(s string) (int, error)`.
- Intermediate: `WrappingErr` — wrap with `%w`, detect with `errors.Is`.
- Advanced: sentinel-error refactor of the same function.

---

### Lesson 6.2 — Goroutines & WaitGroup

**Concept slide**

- `go f()`, `sync.WaitGroup` (Add/Done/Wait).

**Tasks**
- Beginner: 3 goroutines append IDs under a pre-wired mutex.
- Intermediate: `ConcurrentSum(nums []int) int` via goroutines + mutex.
- Advanced: worker-pool `Squares(nums []int) []int` (sorted result).

---

### Lesson 6.3 — Channels & Shared State

**Concept slide**

- `make(chan T)`, `<-`, `close`, `for v := range ch`; buffered vs unbuffered; deadlock gotcha.

**Tasks**
- Beginner: `Producer(ch chan<- int, n int)`.
- Intermediate: `SquaresPipeline(k int) []int` — two goroutine, two-channel pipeline.
- Advanced: `EvenSum(n int) int` — producer + `for range` consumer summing evens.

---

### Lesson 6.4 — Capstone: Concurrent Checkout

**Concept slide (practice only, no new concept)**

- Combine errors + goroutines + channels/order-independence into a tiny checkout app.

**Tasks**
- Beginner: `Validate(item string, qty int) error`.
- Intermediate: worker `go` + WaitGroup posting into a `chan int`.
- Advanced: producer→consumer pipeline with `close(ch)`.

---

## 🧩 Chapter 6 Quiz (5 questions, ~2 min)

1. What built-in type does Go use instead of exceptions? *(the `error` interface)*
2. Which `fmt.Errorf` verb wraps an error? *(`%w`)*
3. What keyword starts a goroutine? *(`go`)*
4. What blocks until all `Done()`s fire? *(`WaitGroup.Wait()`)*
5. What happens receiving from a closed, empty channel? *(zero value + `ok == false`)*