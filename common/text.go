package c

import (
	"encoding/json"
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

func ToJsonString(obj any) string {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

// 把str 變成 12****34
func HideString(str string) string {
	strLength := len(str)
	if strLength <= 6 {
		return str
	}
	iEnd := strLength - 4
	return str[:2] + strings.Repeat("*", iEnd-2) + str[iEnd:]
}
