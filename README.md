# simple-monads

A Rust-inspired implementation of `Option` and `Result` types using generic structs.

Requires: Go >= 1.18 (generics added). 

## Examples
```go
value, err := MyFallibleFunction()
res := Result(value, err)
if res.IsOk(){
    ...
}
```

```go
value, err := MyFallibleFunction()
res := Result(value, err)
value = res.UnwrapOr(value2)
```

Option constructor accepts pointers:
```go
maybe := Option(ReturnsPtr())
if maybe.IsSome(){
    ...
}
```

```go
maybe := Option(ReturnsPtr())
value := maybe.UnwrapOr(myDefaultValue)
```

Use `OptionType` or `ResultType` for type defintions such as function returns:

```go
func MyOptionalInt() OptionType[int]{
    ...
}
```
```go
func MyFallibleInt() ResultType[int]{
    ...
}
```

Supports `Scan` interface from `database/sql` and implements a `ToDB` helper method:
```go
db := Result(sql.Open("sqlite3", ":memory:")).Unwrap()
defer db.Close()

Result(db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, value INTEGER)")).Unwrap()
Result(db.Exec("INSERT INTO test (value) VALUES (1)")).Unwrap()

var i *int64 // sql driver requires int64
opt := Option(i)
_ = db.QueryRow("SELECT value FROM test WHERE id = 1").Scan(&opt)
fmt.Println(opt.Some()) // prints 1
```

```go
db := Result(sql.Open("sqlite3", ":memory:")).Unwrap()
defer db.Close()

Result(db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, value INTEGER)")).Unwrap()

var i *int64 // sql driver requires int64
opt := Option(i)
insert := Result(db.Exec(
	"INSERT INTO test (value) VALUES (?)",
	Result(opt.ToDB()).Unwrap(),
))
```

See tests for a few more examples.
