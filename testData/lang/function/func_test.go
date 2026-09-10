//nolint:testableexamples // This legacy implementation intentionally relies on these patterns.
package function_test

import (
	"fmt"

	"github.com/umats/gomarkdoc/testData/lang/function"
)

func ExampleStandalone() {
	// Comment
	res, _ := function.Standalone(2, "abc")
	fmt.Println(res)
	// Output: 2
}

func ExampleStandalone_zero() {
	res, _ := function.Standalone(0, "def")
	fmt.Println(res)
	// Output: 0
}

func ExampleReceiver() {
	// Add some comments
	r := &function.Receiver{}
	// And some more
	fmt.Println(r)
}

func ExampleReceiver_subTest() {
	var r function.Receiver
	r.WithReceiver()
}

func ExampleGeneric_WithGenericReceiver() {
	r := function.Generic[int]{}
	r.WithGenericReceiver()
}
