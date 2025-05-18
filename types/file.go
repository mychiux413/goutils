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
		s := strings.SplitN(str, "base64,", 2)
		if len(s) != 2 {
			return nil, fmt.Errorf("split base64, but got split length: %d", len(s))
		}
		str = s[1]
	}
	return base64.StdEncoding.DecodeString(str)
}
