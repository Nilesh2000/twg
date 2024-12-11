package example_test

import (
	"fmt"
	"io"
	"twg/example"

	// Needed for side effect
	_ "image/jpeg"
)

var file string = "this is not used"

func Example_crop() {
	var r io.Reader
	img, err := example.Decode(r)
	if err != nil {
		panic(err)
	}

	err = example.Crop(img, 0, 0, 100, 100)
	if err != nil {
		panic(err)
	}

	var w io.Writer
	err = example.Encode(img, w)
	if err != nil {
		panic(err)
	}
	fmt.Println("See out.jpg for the cropped image.")

	// Output:
	// See out.jpg for the cropped image.
}
