// Code generated by optionalize
// on: 2017-10-19 08:41:17.722336045 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type OptionalUInt32Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value UInt32)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (UInt32, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def UInt32) UInt32

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() UInt32) UInt32

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() UInt32

	// And returns None if the optional is None, otherwise returns optb.
	And(optb OptionalUInt32Slice) OptionalUInt32Slice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb OptionalUInt32Slice) OptionalUInt32Slice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() OptionalUInt32Slice) OptionalUInt32Slice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value UInt32)) someOptionalUInt32SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someOptionalUInt32SliceHandler interface {
	None(fn func())
}

type _OptionalUInt32Slice struct {
	hasValue    bool
	unsafeValue UInt32
}

// String conforms to fmt.Stringer interface.
func (opt *_OptionalUInt32Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewOptionalUInt32Slice() OptionalUInt32Slice {
	opt := &_OptionalUInt32Slice{}
	opt.Take()
	return opt
}

func OptionalUInt32SliceFrom(value UInt32, hasValue bool) OptionalUInt32Slice {
	opt := &_OptionalUInt32Slice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_OptionalUInt32Slice) Take() {
	var value UInt32
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_OptionalUInt32Slice) From(value UInt32) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_OptionalUInt32Slice) Unwrap() (UInt32, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_OptionalUInt32Slice) UnwrapOr(def UInt32) UInt32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_OptionalUInt32Slice) UnwrapOrElse(fn func() UInt32) UInt32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_OptionalUInt32Slice) UnwrapOrPanic() UInt32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalUInt32Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_OptionalUInt32Slice) And(optb OptionalUInt32Slice) OptionalUInt32Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_OptionalUInt32Slice) Or(optb OptionalUInt32Slice) OptionalUInt32Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_OptionalUInt32Slice) OrElse(fn func() OptionalUInt32Slice) OptionalUInt32Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_OptionalUInt32Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_OptionalUInt32Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_OptionalUInt32Slice) Some(fn func(value UInt32)) someOptionalUInt32SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someOptionalUInt32SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_OptionalUInt32Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someOptionalUInt32SliceHandler struct {
	opt *_OptionalUInt32Slice
}

func (some _someOptionalUInt32SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_OptionalUInt32Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_OptionalUInt32Slice) getUnsafeValue() UInt32 {
	return opt.unsafeValue
}