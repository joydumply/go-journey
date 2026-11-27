# Go Journey

Personal repo for working through [roadmap.sh](https://roadmap.sh/golang) - learning Go step by step, tracking code and progress along the way.

## Structure

```
.
├── hello/          # first steps, basic syntax
├── maps/           # working with maps
├── slices/         # slices, capacity, conversions
└── ...
```

Each folder is a separate Go module (`go.mod`) with code examples for a specific topic.

## Progress

- [x] Basic syntax
- [x] Arrays
- [x] Slices (capacity, growth, conversions)
- [x] Maps
- [x] Structs 
- [x] Loops
- [x] Functions
- [x] Pointers
- [x] Methods vs Functions
- [x] Pointer Receivers
- [x] Value Receivers
- [x] Interface
- [x] Generics
- [x] Library System Task

## Library System Task

Located in [`22.library/`](22.library/). A small library model that pulls together everything covered so far.

**What it includes:**

- `Author`, `Book` and `Magazine` structs, with `Author` embedded in `Book` and JSON struct tags
- `CheckOut` / `Return` methods with error handling (`errors.New`)
- An `Item` interface implemented by both `Book` and `Magazine`, plus `PrintCatalog` over `[]Item`
- A type switch in `Describe(x any)` that handles values and pointers
- A generic `Queue[T]` with `Enqueue`, `Dequeue` and `Len`
- A generic `Filter[T]` that takes a predicate function
- Stats over a slice of books: `CountByDecade` (map of decade → count), `AgeCategory`, and the variadic `LibraryStats`

**Concepts covered:**

| Topic | Where |
| --- | --- |
| Struct embedding, promoted fields | `Book` embeds `Author` |
| Struct tags, `json.MarshalIndent` | `book.go`, `main.go` |
| Pointer vs value receivers | `CheckOut`/`Return` (pointer), `Describe` (value) |
| Method sets & interface satisfaction | `[]Item{&b, &m1}` |
| Value copy vs pointer in an interface | catalog experiment in `main.go` |
| Error returns and handling | `CheckOut` |
| Type switch on `any` | `Describe(x any)` |
| Generics with type parameters | `Queue[T]`, `Filter[T]` |
| Avoiding memory leaks in slice-based queues | zeroing the dequeued slot |
| Higher-order functions / predicates | `Filter` |
| Maps, `for range`, integer arithmetic | `CountByDecade` |
| Variadic functions, `slice...` | `LibraryStats` |
| Named return values | `LibraryStats` |

## Running

```bash
cd <topic-folder>
go run .
```

## Notes

Detailed notes/write-ups for each topic are kept separately in Obsidian — this repo holds the code only.
