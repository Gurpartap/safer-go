// Code generated by optionalize
// on: 2017-10-19 08:41:19.590930488 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type OptionalFloat32Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value Float32)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (Float32, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def Float32) Float32

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() Float32) Float32

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() Float32

	// And returns None if the optional is None, otherwise returns optb.
	And(optb OptionalFloat32Slice) OptionalFloat32Slice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb OptionalFloat32Slice) OptionalFloat32Slice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() OptionalFloat32Slice) OptionalFloat32Slice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value Float32)) someOptionalFloat32SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someOptionalFloat32SliceHandler interface {
	None(fn func())
}

type _OptionalFloat32Slice struct {
	hasValue    bool
	unsafeValue Float32
}

// String conforms to fmt.Stringer interface.
func (opt *_OptionalFloat32Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewOptionalFloat32Slice() OptionalFloat32Slice {
	opt := &_OptionalFloat32Slice{}
	opt.Take()
	return opt
}

func OptionalFloat32SliceFrom(value Float32, hasValue bool) OptionalFloat32Slice {
	opt := &_OptionalFloat32Slice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_OptionalFloat32Slice) Take() {
	var value Float32
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_OptionalFloat32Slice) From(value Float32) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_OptionalFloat32Slice) Unwrap() (Float32, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_OptionalFloat32Slice) UnwrapOr(def Float32) Float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_OptionalFloat32Slice) UnwrapOrElse(fn func() Float32) Float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_OptionalFloat32Slice) UnwrapOrPanic() Float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalFloat32Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_OptionalFloat32Slice) And(optb OptionalFloat32Slice) OptionalFloat32Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_OptionalFloat32Slice) Or(optb OptionalFloat32Slice) OptionalFloat32Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_OptionalFloat32Slice) OrElse(fn func() OptionalFloat32Slice) OptionalFloat32Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_OptionalFloat32Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_OptionalFloat32Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_OptionalFloat32Slice) Some(fn func(value Float32)) someOptionalFloat32SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someOptionalFloat32SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_OptionalFloat32Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someOptionalFloat32SliceHandler struct {
	opt *_OptionalFloat32Slice
}

func (some _someOptionalFloat32SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_OptionalFloat32Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_OptionalFloat32Slice) getUnsafeValue() Float32 {
	return opt.unsafeValue
}
