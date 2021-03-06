// Code generated by optionalize
// on: 2017-10-19 12:55:49.544189192 &#43;0000 UTC
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

type Bool struct {
	hasValue    bool
	unsafeValue bool
}

func NewBool() Bool {
	opt := &Bool{}
	opt.Take()
	return *opt
}

func BoolFrom(value bool, hasValue bool) Bool {
	opt := &Bool{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *Bool) Take() {
	var value bool
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *Bool) From(value bool) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *Bool) Unwrap() (bool, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *Bool) UnwrapOr(def bool) bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *Bool) UnwrapOrElse(fn func() bool) bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *Bool) UnwrapOrPanic() bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Bool")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *Bool) And(optb Bool) Bool {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *Bool) Or(optb Bool) Bool {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *Bool) OrElse(fn func() Bool) Bool {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *Bool) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *Bool) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *Bool) Some(fn func(value bool)) someBoolHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someBoolHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *Bool) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someBoolHandler struct {
	opt *Bool
}

func (some someBoolHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *Bool) getHasValue() bool {
	return opt.hasValue
}

func (opt *Bool) getUnsafeValue() bool {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *Bool) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt Bool) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *Bool) UnmarshalJSON(data []byte) error {
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
func (opt *Bool) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue bool
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt Bool) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
