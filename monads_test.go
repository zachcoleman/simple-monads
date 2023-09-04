package monads

import (
	"testing"
)

func TestOption(t *testing.T) {
	var i *int
	o := Option(i)
	if !o.IsNone() {
		t.Error()
	}
	if o.IsSome() {
		t.Error()
	}
	_ = o.None()
	if o.UnwrapOr(1) != 1 {
		t.Error()
	}
	if o.UnwrapOrElse(func() int { return 1 }) != 1 {
		t.Error()
	}

	var j int = 1
	o = Option(&j)
	if o.IsNone() {
		t.Error()
	}
	if !o.IsSome() {
		t.Error()
	}
	if o.Some() != 1 {
		t.Error()
	}
	if o.Unwrap() != 1 {
		t.Error()
	}
	if o.UnwrapOr(2) != 1 {
		t.Error()
	}
	if o.UnwrapOrElse(func() int { return 2 }) != 1 {
		t.Error()
	}
}

type Err string

func (e Err) Error() string { return string(e) }
func Fallible(shouldFail bool) (int, error) {
	if shouldFail {
		return 0, Err("fail")
	}
	return 1, nil
}

func TestResult(t *testing.T) {
	i, err := Fallible(false)
	r := Result(i, err)
	if r.IsErr() {
		t.Error()
	}
	if !r.IsOk() {
		t.Error()
	}
	_ = r.Ok()
	if r.Unwrap() != 1 {
		t.Error()
	}
	if r.UnwrapOr(2) != 1 {
		t.Error()
	}
	if r.UnwrapOrElse(func() int { return 2 }) != 1 {
		t.Error()
	}

	i, err = Fallible(true)
	r = Result(i, err)
	if !r.IsErr() {
		t.Error()
	}
	if r.IsOk() {
		t.Error()
	}
	_ = r.Err()
	if r.UnwrapOr(2) != 2 {
		t.Error()
	}
	if r.UnwrapOrElse(func() int { return 2 }) != 2 {
		t.Error()
	}
}

func TestOptionPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error()
		}
	}()
	var i *int
	o := Option(i)
	_ = o.Some()
}

func TestResultPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error()
		}
	}()
	i, err := Fallible(false)
	r := Result(i, err)
	_ = r.Err()
}

func TestFuncReturnType(t *testing.T) {
	_ = func(v float64) OptionType[int] {
		if v > 0.5 {
			var k int = 1
			return Option(&k)
		} else {
			return Option[int](nil)
		}
	}
	_ = func(v float64) OptionType[int] {
		if v > 0.5 {
			var k int = 1
			return Option(&k)
		} else {
			var k *int
			return Option(k)
		}
	}
	_ = func(v float64) ResultType[int] {
		if v > 0.5 {
			return Result(0, nil)
		}
		return Result(0, Err("fail"))
	}
}
