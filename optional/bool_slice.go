// Code generated by optionalize
// on: 2017-10-19 12:55:55.372824079 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/Gurpartap/safer-go/optional/internal/convert"
	"github.com/pkg/errors"
)

type BoolSlice struct {
	hasValue    bool
	unsafeValue []bool
}

func NewBoolSlice() BoolSlice {
	opt := &BoolSlice{}
	opt.Take()
	return *opt
}

func BoolSliceFrom(value []bool, hasValue bool) BoolSlice {
	opt := &BoolSlice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *BoolSlice) Take() {
	var value []bool
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *BoolSlice) From(value []bool) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *BoolSlice) Unwrap() ([]bool, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *BoolSlice) UnwrapOr(def []bool) []bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *BoolSlice) UnwrapOrElse(fn func() []bool) []bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *BoolSlice) UnwrapOrPanic() []bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap BoolSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *BoolSlice) And(optb BoolSlice) BoolSlice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *BoolSlice) Or(optb BoolSlice) BoolSlice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *BoolSlice) OrElse(fn func() BoolSlice) BoolSlice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *BoolSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *BoolSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *BoolSlice) Some(fn func(value []bool)) someBoolSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someBoolSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *BoolSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someBoolSliceHandler struct {
	opt *BoolSlice
}

func (some someBoolSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *BoolSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *BoolSlice) getUnsafeValue() []bool {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *BoolSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt BoolSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *BoolSlice) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) || data == nil {
		opt.Take()
		return nil
	}

	err := json.Unmarshal(data, &opt.unsafeValue)
	if err != nil {
		opt.hasValue = false
		return errors.WithStack(err)
	}
	opt.hasValue = true

	return nil
}

// Scan implements the sql Scanner interface.
func (opt *BoolSlice) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue []bool
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt BoolSlice) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
