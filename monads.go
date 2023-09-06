package monads

import (
	"database/sql/driver"
	"errors"
	"time"
)

type SQLDriverCompatible interface {
	Scan(src any) error
	Value() (driver.Value, error)
}

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
func (o *OptionType[T]) IsNone() bool { return o.none }
func (o *OptionType[T]) IsSome() bool { return !o.none }
func (o *OptionType[T]) Some() T {
	if o.none {
		panic("some on none")
	}
	return o.some
}
func (o *OptionType[T]) None() *T {
	if !o.none {
		panic("none on some")
	}
	return nil
}
func (o *OptionType[T]) Unwrap() T {
	if o.none {
		panic("unwrap on none")
	}
	return o.some
}
func (o *OptionType[T]) UnwrapOr(def T) T {
	if o.none {
		return def
	}
	return o.some
}
func (o *OptionType[T]) UnwrapOrElse(f func() T) T {
	if o.none {
		return f()
	}
	return o.some
}

func (o *OptionType[T]) Scan(src any) error {
	if src == nil {
		o.none = true
		return nil
	}
	switch src.(type) {
	case int64, float64, bool, []byte, string, time.Time:
		// check for sql driver compatibility w/ recieving type
		switch any(o.some).(type) {
		case int, int32, float32:
			return errors.New("incompatible types w/ sql driver")
		default:
			o.some = src.(T)
			o.none = false
			return nil
		}
	default:
		return any(o.some).(SQLDriverCompatible).Scan(src)
	}
}

func (o *OptionType[T]) ToDB() (driver.Value, error) {
	if o.IsNone() {
		return nil, nil
	}

	switch any(o.some).(type) {
	case int, int32, int64, float32, float64, bool, []byte, string, time.Time:
		return o.some, nil
	default:
		return any(o.some).(SQLDriverCompatible).Value()
	}
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
