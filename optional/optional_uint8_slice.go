// Code generated by optionalize
// on: 2017-09-10 10:30:44.596534468 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type OptionalUInt8Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value UInt8)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (UInt8, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def UInt8) UInt8

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() UInt8) UInt8

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() UInt8

	// And returns None if the optional is None, otherwise returns optb.
	And(optb OptionalUInt8Slice) OptionalUInt8Slice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb OptionalUInt8Slice) OptionalUInt8Slice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() OptionalUInt8Slice) OptionalUInt8Slice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value UInt8)) someOptionalUInt8SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someOptionalUInt8SliceHandler interface {
	None(fn func())
}

type _OptionalUInt8Slice struct {
	hasValue    bool
	unsafeValue UInt8
}

// String conforms to fmt.Stringer interface.
func (opt *_OptionalUInt8Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewOptionalUInt8Slice() OptionalUInt8Slice {
	opt := &_OptionalUInt8Slice{}
	opt.Take()
	return opt
}

func OptionalUInt8SliceFrom(value UInt8) OptionalUInt8Slice {
	opt := &_OptionalUInt8Slice{}
	opt.From(value)
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_OptionalUInt8Slice) Take() {
	var value UInt8
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_OptionalUInt8Slice) From(value UInt8) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_OptionalUInt8Slice) Unwrap() (UInt8, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_OptionalUInt8Slice) UnwrapOr(def UInt8) UInt8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_OptionalUInt8Slice) UnwrapOrElse(fn func() UInt8) UInt8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_OptionalUInt8Slice) UnwrapOrPanic() UInt8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalUInt8Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_OptionalUInt8Slice) And(optb OptionalUInt8Slice) OptionalUInt8Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_OptionalUInt8Slice) Or(optb OptionalUInt8Slice) OptionalUInt8Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_OptionalUInt8Slice) OrElse(fn func() OptionalUInt8Slice) OptionalUInt8Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_OptionalUInt8Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_OptionalUInt8Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_OptionalUInt8Slice) Some(fn func(value UInt8)) someOptionalUInt8SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someOptionalUInt8SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_OptionalUInt8Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someOptionalUInt8SliceHandler struct {
	opt *_OptionalUInt8Slice
}

func (some _someOptionalUInt8SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_OptionalUInt8Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_OptionalUInt8Slice) getUnsafeValue() UInt8 {
	return opt.unsafeValue
}
