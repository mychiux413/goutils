package c_test

import (
	"fmt"

	c "github.com/mychiux413/goutils/common"
)

func ExampleChangeFilenameExt() {
	outputFilename, err := c.ChangeFilenameExt("filename.webp", ".png")
	if err != nil {
		fmt.Printf("ChangeFilenameExt got error: %v\n", err)
	}
	fmt.Println(outputFilename)
	_, err = c.ChangeFilenameExt("filename.webp", ".webp")
	fmt.Println(err)
	// Output:
	// filename.png
	// .webp == .webp
}
