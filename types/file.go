package t

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

type FileData = []byte

func MarshalFileData(src []byte) graphql.Marshaler {
	fmt.Println("Marshal File")
	output := base64.StdEncoding.EncodeToString(src)
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(output))
	})
}

func UnmarshalFileData(v interface{}) ([]byte, error) {
	fmt.Println("Unmarshal File")
	str, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("file must be strings in base64")
	}
	if len(str) > 30 && strings.Contains(str[:30], "base64,") {
		// 相容 data:image/jpeg;base64,<BASE64> 的版本
		str = strings.Split(str, "base64,")[1]
	}
	return base64.StdEncoding.DecodeString(str)
}
