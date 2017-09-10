// Code generated by optionalize
// on: 2017-09-10 10:30:39.41458165 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type UInt64Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []uint64)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]uint64, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []uint64) []uint64

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []uint64) []uint64

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []uint64

	// And returns None if the optional is None, otherwise returns optb.
	And(optb UInt64Slice) UInt64Slice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb UInt64Slice) UInt64Slice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() UInt64Slice) UInt64Slice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []uint64)) someUInt64SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someUInt64SliceHandler interface {
	None(fn func())
}

type _UInt64Slice struct {
	hasValue    bool
	unsafeValue []uint64
}

// String conforms to fmt.Stringer interface.
func (opt *_UInt64Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewUInt64Slice() UInt64Slice {
	opt := &_UInt64Slice{}
	opt.Take()
	return opt
}

func UInt64SliceFrom(value []uint64) UInt64Slice {
	opt := &_UInt64Slice{}
	opt.From(value)
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_UInt64Slice) Take() {
	var value []uint64
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_UInt64Slice) From(value []uint64) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_UInt64Slice) Unwrap() ([]uint64, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_UInt64Slice) UnwrapOr(def []uint64) []uint64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_UInt64Slice) UnwrapOrElse(fn func() []uint64) []uint64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_UInt64Slice) UnwrapOrPanic() []uint64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap UInt64Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_UInt64Slice) And(optb UInt64Slice) UInt64Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_UInt64Slice) Or(optb UInt64Slice) UInt64Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_UInt64Slice) OrElse(fn func() UInt64Slice) UInt64Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_UInt64Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_UInt64Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_UInt64Slice) Some(fn func(value []uint64)) someUInt64SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someUInt64SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_UInt64Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someUInt64SliceHandler struct {
	opt *_UInt64Slice
}

func (some _someUInt64SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_UInt64Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_UInt64Slice) getUnsafeValue() []uint64 {
	return opt.unsafeValue
}
