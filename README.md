# simple-monads

A Rust-inspired implementation of `Option` and `Result` types using generic structs.

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

func MyFallibleInt() ResultType[int]{
    ...
}
```

See tests for a few more examples.