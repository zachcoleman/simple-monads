package monads

type OptionType[T any] struct {
	some T
	none bool
}

func Option[T any](v *T) OptionType[T] {
	if v == nil {
		return OptionType[T]{none: true}
	}
	return OptionType[T]{
		some: *v,
		none: false,
	}
}
func (o OptionType[T]) IsNone() bool { return o.none }
func (o OptionType[T]) IsSome() bool { return !o.none }
func (o OptionType[T]) Some() T {
	if o.none {
		panic("some on none")
	}
	return o.some
}
func (o OptionType[T]) None() *T {
	if !o.none {
		panic("none on some")
	}
	return nil
}
func (o OptionType[T]) Unwrap() T {
	if o.none {
		panic("unwrap on none")
	}
	return o.some
}
func (o OptionType[T]) UnwrapOr(def T) T {
	if o.none {
		return def
	}
	return o.some
}
func (o OptionType[T]) UnwrapOrElse(f func() T) T {
	if o.none {
		return f()
	}
	return o.some
}

type ResultType[T any] struct {
	ok  T
	err error
}

func Result[T any](v T, e error) ResultType[T] {
	return ResultType[T]{ok: v, err: e}
}

func (r ResultType[T]) IsOk() bool  { return r.err == nil }
func (r ResultType[T]) IsErr() bool { return r.err != nil }
func (r ResultType[T]) Ok() T {
	if r.err != nil {
		panic("ok on err")
	}
	return r.ok
}
func (r ResultType[T]) Err() error {
	if r.err == nil {
		panic("err on ok")
	}
	return r.err
}
func (r ResultType[T]) Unwrap() T {
	if r.err != nil {
		panic("unwrap on err")
	}
	return r.ok
}
func (r ResultType[T]) UnwrapOr(def T) T {
	if r.err != nil {
		return def
	}
	return r.ok
}
func (r ResultType[T]) UnwrapOrElse(f func() T) T {
	if r.err != nil {
		return f()
	}
	return r.ok
}
