package encoding_test

import (
	"testing"
	"unsafe"

	"github.com/alxarno/parquet-go-athena-list-hack/encoding"
)

func TestValuesSize(t *testing.T) {
	t.Log(unsafe.Sizeof(encoding.Values{}))
}
