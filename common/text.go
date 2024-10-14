package c

import "strings"

// 把所有的string都移除掉空白
func MaybeTrimSpace(strPtrs ...*string) {
	for _, strPtr := range strPtrs {
		if strPtr == nil {
			continue
		}
		*strPtr = strings.TrimSpace(*strPtr)
	}
}
