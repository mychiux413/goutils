package c

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ChangeFilenameExt(filename string, toExtWithDot string) (string, error) {
	if strings.Contains(filename, "/") || !strings.Contains(filename, ".") {
		return "", fmt.Errorf("invalid filename: %s", filename)
	}
	if !strings.HasPrefix(toExtWithDot, ".") {
		toExtWithDot = "." + toExtWithDot
	}
	beforeWithDot := filepath.Ext(filename)
	if beforeWithDot == toExtWithDot {
		return "", fmt.Errorf("%s == %s", beforeWithDot, toExtWithDot)
	}
	return strings.Replace(filename, beforeWithDot, toExtWithDot, 1), nil
}
