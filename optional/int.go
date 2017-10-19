// Code generated by optionalize
// on: 2017-10-19 08:58:18.468746801 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Int interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value int)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (int, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def int) int

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() int) int

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() int

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Int) Int

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Int) Int

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Int) Int

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value int)) someIntHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someIntHandler interface {
	None(fn func())
}

type _Int struct {
	hasValue    bool
	unsafeValue int
}

func NewInt() Int {
	opt := &_Int{}
	opt.Take()
	return opt
}

func IntFrom(value int, hasValue bool) Int {
	opt := &_Int{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Int) Take() {
	var value int
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Int) From(value int) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Int) Unwrap() (int, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Int) UnwrapOr(def int) int {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Int) UnwrapOrElse(fn func() int) int {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Int) UnwrapOrPanic() int {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Int")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Int) And(optb Int) Int {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Int) Or(optb Int) Int {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Int) OrElse(fn func() Int) Int {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Int) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Int) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Int) Some(fn func(value int)) someIntHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someIntHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Int) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someIntHandler struct {
	opt *_Int
}

func (some _someIntHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Int) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Int) getUnsafeValue() int {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_Int) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}
