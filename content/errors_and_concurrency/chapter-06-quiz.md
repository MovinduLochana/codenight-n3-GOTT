# Chapter 6 Quiz — Errors & a Taste of Concurrency

Five quick questions (~2 min). Answers are shown on the next screen after the student responds.

---

**Q1. What built-in type does Go use instead of exceptions?**

- A) `panic` values
- B) The `error` interface
- C) `nil` booleans

**Answer:** B — Go functions return an `error` as a normal value; there are no exceptions.

---

**Q2. Which `fmt.Errorf` verb wraps an error?**

- A) `%v`
- B) `%e`
- C) `%w`

**Answer:** C — `%w` wraps the error and preserves the chain so `errors.Is`/`errors.As` can inspect it.

---

**Q3. What keyword starts a goroutine?**

- A) `thread`
- B) `go`
- C) `async`

**Answer:** B — `go someFunc()` launches `someFunc` as a lightweight concurrent goroutine.

---

**Q4. What blocks until all `Done()`s fire?**

- A) `wg.Wait()`
- B) `wg.Stop()`
- C) `close(wg)`

**Answer:** A — `sync.WaitGroup.Wait()` blocks until every `Add`'d goroutine has called `Done()`.

---

**Q5. What happens receiving from a closed, empty channel?**

- A) It panics
- B) It returns the zero value, with `ok == false`
- C) It blocks forever

**Answer:** B — a `for range` over a closed, drained channel just stops; a single receive returns the zero value and `ok == false`.