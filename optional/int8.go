// Code generated by optionalize
// on: 2017-09-10 10:30:33.594644222 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Int8 interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value int8)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (int8, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def int8) int8

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() int8) int8

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() int8

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Int8) Int8

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb Int8) Int8

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() Int8) Int8

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value int8)) someInt8Handler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someInt8Handler interface {
	None(fn func())
}

type _Int8 struct {
	hasValue    bool
	unsafeValue int8
}

// String conforms to fmt.Stringer interface.
func (opt *_Int8) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewInt8() Int8 {
	opt := &_Int8{}
	opt.Take()
	return opt
}

func Int8From(value int8) Int8 {
	opt := &_Int8{}
	opt.From(value)
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_Int8) Take() {
	var value int8
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Int8) From(value int8) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Int8) Unwrap() (int8, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Int8) UnwrapOr(def int8) int8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Int8) UnwrapOrElse(fn func() int8) int8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Int8) UnwrapOrPanic() int8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Int8")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Int8) And(optb Int8) Int8 {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_Int8) Or(optb Int8) Int8 {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Int8) OrElse(fn func() Int8) Int8 {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_Int8) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_Int8) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Int8) Some(fn func(value int8)) someInt8Handler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someInt8Handler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Int8) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someInt8Handler struct {
	opt *_Int8
}

func (some _someInt8Handler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Int8) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Int8) getUnsafeValue() int8 {
	return opt.unsafeValue
}
