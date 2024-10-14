package c

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/andybalholm/brotli"
)

/*
發出http請求時，帶入Header: "br, gzip, deflate"

例如：

	http.Header{
		"Accept-Encoding": []string{"br", "gzip", "deflate"},
	}

code block:

	response, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}
	reader, err := common.BrGzipDecompressor(response)
	if err != nil {
		panic(err)
	}
*/
func BrGzipDecompressor(response *http.Response) (io.Reader, error) {
	var reader io.Reader = response.Body
	if !response.Uncompressed {
		contentEncoding := response.Header.Get("Content-Encoding")
		switch contentEncoding {
		case "", "deflate":
		case "br":
			if DEBUG {
				fmt.Println("----> Enable br decompressor")
			}
			reader = brotli.NewReader(response.Body)
			response.Uncompressed = true
		case "gzip":
			if DEBUG {
				fmt.Println("----> Enable gzip decompressor")
			}
			r, err := gzip.NewReader(response.Body)
			if err != nil {
				return nil, err
			}
			reader = r
			response.Uncompressed = true
		default:
			log.Printf("[skip error] use deflate decode for unknown content-encoding: %s\n", contentEncoding)
		}
	}
	return reader, nil
}
