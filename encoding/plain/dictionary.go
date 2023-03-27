package plain

import (
	"github.com/alxarno/parquet-go-athena-list-hack/encoding"
	"github.com/alxarno/parquet-go-athena-list-hack/format"
)

type DictionaryEncoding struct {
	encoding.NotSupported
	plain Encoding
}

func (e *DictionaryEncoding) String() string {
	return "PLAIN_DICTIONARY"
}

func (e *DictionaryEncoding) Encoding() format.Encoding {
	return format.PlainDictionary
}

func (e *DictionaryEncoding) EncodeInt32(dst []byte, src []int32) ([]byte, error) {
	return e.plain.EncodeInt32(dst, src)
}

func (e *DictionaryEncoding) DecodeInt32(dst []int32, src []byte) ([]int32, error) {
	return e.plain.DecodeInt32(dst, src)
}
