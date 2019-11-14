package bitutil

import (
	"fmt"
	"testing"
)

func TestExample_CompressEncode(t *testing.T) {
	bit := []byte("hello")
	compress := CompressBytes(bit)
	fmt.Println(string(compress))

	bit = append([]byte{'h', 0, 'e', 0, 0, 1, 0}, "llo"...)
	compress = CompressBytes(bit)
	fmt.Println(string(compress))
}
