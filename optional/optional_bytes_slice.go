// Code generated by optionalize
// on: 2017-10-19 08:41:21.793767983 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type OptionalBytesSlice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value Bytes)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (Bytes, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def Bytes) Bytes

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() Bytes) Bytes

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() Bytes

	// And returns None if the optional is None, otherwise returns optb.
	And(optb OptionalBytesSlice) OptionalBytesSlice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb OptionalBytesSlice) OptionalBytesSlice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() OptionalBytesSlice) OptionalBytesSlice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value Bytes)) someOptionalBytesSliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someOptionalBytesSliceHandler interface {
	None(fn func())
}

type _OptionalBytesSlice struct {
	hasValue    bool
	unsafeValue Bytes
}

// String conforms to fmt.Stringer interface.
func (opt *_OptionalBytesSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewOptionalBytesSlice() OptionalBytesSlice {
	opt := &_OptionalBytesSlice{}
	opt.Take()
	return opt
}

func OptionalBytesSliceFrom(value Bytes, hasValue bool) OptionalBytesSlice {
	opt := &_OptionalBytesSlice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_OptionalBytesSlice) Take() {
	var value Bytes
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_OptionalBytesSlice) From(value Bytes) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_OptionalBytesSlice) Unwrap() (Bytes, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_OptionalBytesSlice) UnwrapOr(def Bytes) Bytes {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_OptionalBytesSlice) UnwrapOrElse(fn func() Bytes) Bytes {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_OptionalBytesSlice) UnwrapOrPanic() Bytes {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalBytesSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_OptionalBytesSlice) And(optb OptionalBytesSlice) OptionalBytesSlice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_OptionalBytesSlice) Or(optb OptionalBytesSlice) OptionalBytesSlice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_OptionalBytesSlice) OrElse(fn func() OptionalBytesSlice) OptionalBytesSlice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_OptionalBytesSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_OptionalBytesSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_OptionalBytesSlice) Some(fn func(value Bytes)) someOptionalBytesSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someOptionalBytesSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_OptionalBytesSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someOptionalBytesSliceHandler struct {
	opt *_OptionalBytesSlice
}

func (some _someOptionalBytesSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_OptionalBytesSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_OptionalBytesSlice) getUnsafeValue() Bytes {
	return opt.unsafeValue
}
