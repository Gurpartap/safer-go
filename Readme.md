# Safer Go types (optionals et al.)

[![GoDoc](https://godoc.org/github.com/Gurpartap/safer-go?status.svg)](https://godoc.org/github.com/Gurpartap/safer-go)

We don't have something like `Optional<T>` _built-in_ type in Go, and I don't always want to use pointers to declare nullable variables.

I'm not even asking for [Generics in Go 2](https://twitter.com/Gurpartap/status/898660498874408960).

## Generate Your `Optional<T>`

```bash
$ cd $GOPATH/github.com/Gurpartap/safer-go/optional/
$ go run cmd/optionalize/main.go -wrapper=Int64 -wrapped=int64
```

Result:

```bash
$ cat int64.go
```

```go
package optional

type Int64 interface {
    // Take takes the value out of the option, leaving a None in its place.
    Take()

    // From performs the conversion.
    From(value int64)

    // Unwrap moves the value out of the optional, if it is Some(value).
    // This function returns multiple values, and if that's undesirable,
    // consider using Some and None functions.
    Unwrap() (int64, bool)

    // UnwrapOr returns the contained value or a default.
    UnwrapOr(def int64) int64

    // UnwrapOrElse returns the contained value or computes it from a closure.
    UnwrapOrElse(fn func() int64) int64

    // UnwrapOrPanic returns the contained value or panics.
    UnwrapOrPanic() int64

    // And returns None if the optional is None, otherwise returns optb.
    And(optb Int64) Int64

    // Or returns the option if it contains a value, otherwise returns optb.
    Or(optb Int64) Int64

    // OrElse returns the option if it contains a value, otherwise calls fn and
    // returns the result.
    OrElse(fn func() Int64) Int64

    // IsSome returns true if the option is a Some value.
    IsSome() bool

    // IsNone returns true if the option is a None value.
    IsNone() bool

    // Some executes the given closure if there is a Some value.
    Some(fn func(value int64)) someInt64Handler

    // None executes the given closure if there is a None value.
    None(fn func())
}

type someInt64Handler interface {
    None(fn func())
}

func NewInt64() Int64 {
    // ...
}

func Int64From(value int64) Int64 {
    // ...
}

// <Int64 Implementation />
```

Replace `Int64` and `int64` with your wrapper type name and wrapped type respectively. Example:

```bash
# not necessary for the generator but assuming we have:
$ cat address.go
package models

type Address struct {
    street  string
    city    string
    country string
}
```

```bash
$ go run cmd/optionalize/main.go -wrapper=OptionalAddress -wrapped=Address
```

## Built-In Optional Types

| wrapped | optional \<T> | optional (slice of \<T>) | optional (slice of (optional \<T>)) |
|---|---|---|---|
| bool | optional.Bool | optional.BoolSlice | optional.OptionalBoolSlice |
| uint8 | optional.UInt8 | optional.UInt8Slice | optional.OptionalUInt8Slice |
| uint16 | optional.UInt16 | optional.UInt16Slice | optional.OptionalUInt16Slice |
| uint32 | optional.UInt32 | optional.UInt32Slice | optional.OptionalUInt32Slice |
| uint64 | optional.UInt64 | optional.UInt64Slice | optional.OptionalUInt64Slice |
| int8 | optional.Int8 | optional.Int8Slice | optional.OptionalInt8Slice |
| int16 | optional.Int16 | optional.Int16Slice | optional.OptionalInt16Slice |
| int32 | optional.Int32 | optional.Int32Slice | optional.OptionalInt32Slice |
| int64 | optional.Int64 | optional.Int64Slice | optional.OptionalInt64Slice |
| float32 | optional.Float32 | optional.Float32Slice | optional.OptionalFloat32Slice |
| float64 | optional.Float64 | optional.Float64Slice | optional.OptionalFloat64Slice |
| complex64 | optional.Complex64 | optional.Complex64Slice | optional.OptionalComplex64Slice |
| complex128 | optional.Complex128 | optional.Complex128Slice | optional.OptionalComplex128Slice |
| string | optional.String | optional.StringSlice | optional.OptionalStringSlice |
| int | optional.Int | optional.IntSlice | optional.OptionalIntSlice |
| uint | optional.UInt | optional.UIntSlice | optional.OptionalUIntSlice |
| bytes | optional.Bytes | optional.BytesSlice | optional.OptionalBytesSlice |
| rune | optional.Rune | optional.RuneSlice | optional.OptionalRuneSlice |
| error | optional.Error | optional.ErrorSlice | optional.OptionalErrorSlice |
| time | optional.Time | optional.TimeSlice | optional.OptionalTimeSlice |


## Usage

```go
import "github.com/Gurpartap/safer-go/optional"
```

```go
var opta optional.Int64              // optional <T>
var optb optional.Int64Slice         // optional (slice of <T>)
var optc optional.OptionalInt64Slice // optional (slice of (optional <T>))
```

### Example

```go
// opt := optional.NewString()
// opt.From("foo")

opt := optional.StringFrom("foo")

fmt.Println("opt:", opt)                         // => opt: Optional("foo")

// 1
if opt, ok := opt.Unwrap(); ok {
    // some
    fmt.Println("opt:", opt)                     // => opt: foo
} else {
    // none
    fmt.Println("opt not set")
}

// 2
opt.
    Some(func(opt string) {
        fmt.Println("opt:", opt)                 // => opt: foo
    }).
    None(func() {
        fmt.Println("opt not set")
    })

// 3
if opt.IsSome() {
    opt := opt.UnwrapOrPanic()
    fmt.Println("opt:", opt)                     // => opt: foo
}
if opt.IsNone() {
    opt := opt.UnwrapOrElse(func() string {
        return "abc"
    })
    fmt.Printf("opt not set. using: %s\n", opt)
}
```

## Credits

This package is inspired by, and inherits bits of code and experience from each of the following:

- https://github.com/piotrkowalczuk/ntypes
- https://github.com/leighmcculloch/go-optional
- https://github.com/markphelps/optional
- https://doc.rust-lang.org/std/option/
- https://doc.rust-lang.org/std/option/enum.Option.html
- https://github.com/apple/swift/blob/master/stdlib/public/core/Optional.swift

## License

Copyright 2017 Gurpartap Singh

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.