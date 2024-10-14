package c_test

import (
	"fmt"

	c "github.com/mychiux413/goutils/common"
)

func ExampleChangeFilenameExt() {
	fmt.Println(c.ChangeFilenameExt("filename.webp", ".png"))
	fmt.Println(c.ChangeFilenameExt("filename.webp", ".webp"))
	// Output:
	// filename.png
	// filename.webp
}
