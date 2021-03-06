// Code generated by optionalize
// on: 2017-10-19 12:56:00.31910022 &#43;0000 UTC
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

type BytesSlice struct {
	hasValue    bool
	unsafeValue [][]byte
}

func NewBytesSlice() BytesSlice {
	opt := &BytesSlice{}
	opt.Take()
	return *opt
}

func BytesSliceFrom(value [][]byte, hasValue bool) BytesSlice {
	opt := &BytesSlice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *BytesSlice) Take() {
	var value [][]byte
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *BytesSlice) From(value [][]byte) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *BytesSlice) Unwrap() ([][]byte, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *BytesSlice) UnwrapOr(def [][]byte) [][]byte {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *BytesSlice) UnwrapOrElse(fn func() [][]byte) [][]byte {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *BytesSlice) UnwrapOrPanic() [][]byte {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap BytesSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *BytesSlice) And(optb BytesSlice) BytesSlice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *BytesSlice) Or(optb BytesSlice) BytesSlice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *BytesSlice) OrElse(fn func() BytesSlice) BytesSlice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *BytesSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *BytesSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *BytesSlice) Some(fn func(value [][]byte)) someBytesSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someBytesSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *BytesSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someBytesSliceHandler struct {
	opt *BytesSlice
}

func (some someBytesSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *BytesSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *BytesSlice) getUnsafeValue() [][]byte {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *BytesSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt BytesSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *BytesSlice) UnmarshalJSON(data []byte) error {
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
func (opt *BytesSlice) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue [][]byte
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt BytesSlice) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
