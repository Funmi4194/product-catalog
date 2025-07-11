# 🛒 Product Catalog Sorter (Go)

An extensible product sorting solution built in **Go**, designed for dynamic ordering based on different business goals (e.g., price, conversion, recency).

> This project demonstrates my ability to design sclable systemsbusing clean code, interfaces, unit testing, and software design pattern that is ideal for backend engineering roles.

---

## Worked on Features

- ✅ Sort by **price** or **conversion rate** (sales/views)
- ✅ Follows the **Strategy Pattern** for easy extension
- ✅ Built with **interface-driven design** for testability
- ✅ Unit-tested with **100% coverage**
- ✅ Uses **table-driven tests** and **Testify** for clean validation
- ✅ Clear project architecture (no import cycles)

---


---

## 🧠 Design Principles

### Strategy Pattern

Each sorting algorithm is implemented as its own type that satisfies the `SortType` interface:

```go
type SortType interface {
    Sort([]Product) []Product
}

You can easily create a new sorter:
type ByCreated struct{}
func (s ByCreated) Sort(products []Product) []Product { ... }
