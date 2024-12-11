package example

import (
	"fmt"
)

func ExampleDemo_Hello() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Tom":   false,
		"Jon":   false,
		"Jane":  true,
	}
	Page(checkIns)

	// Unordered Output:
	// Paging Alice; please see the front desk to check in.
	// Paging Jon; please see the front desk to check in.
	// Paging Tom; please see the front desk to check in.
}
