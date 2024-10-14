package c

import "os"

var DEBUG bool = os.Getenv("DEBUG") != ""
