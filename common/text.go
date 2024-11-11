package c

import (
	"regexp"
	"strings"
)

var abnormalCharReg *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9]`)

// 把所有的string都移除掉空白
func MaybeTrimSpace(strPtrs ...*string) {
	for _, strPtr := range strPtrs {
		if strPtr == nil {
			continue
		}
		*strPtr = strings.TrimSpace(*strPtr)
	}
}

// 把特殊字元清除掉
func ClearAbnormalChars(str string) string {
	output := abnormalCharReg.ReplaceAllString(str, "")
	return output
}
