package main

import (
	"fmt"

	"github.com/Gurpartap/safer-go/optional"
)

func main() {
	opt := optional.NewString()
	opt.From("foo")

	fmt.Println("opt:", opt)

	if opt, ok := opt.Unwrap(); ok {
		// some
		fmt.Println("opt:", opt)
	} else {
		// none
		fmt.Println("opt not set")
	}

	opt.
		Some(func(opt string) {
			fmt.Println("opt:", opt)
		}).
		None(func() {
			fmt.Println("opt not set")
		})

	if opt.IsSome() {
		opt := opt.UnwrapOrPanic()
		fmt.Println("opt:", opt)
	}

	if opt.IsNone() {
		opt := opt.UnwrapOrElse(func() string {
			return "abc"
		})
		fmt.Printf("opt not set. using: %s\n", opt)
	}
}
