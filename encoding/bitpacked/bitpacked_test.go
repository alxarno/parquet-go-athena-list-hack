//go:build go1.18
// +build go1.18

package bitpacked_test

import (
	"testing"

	"github.com/alxarno/parquet-go-athena-list-hack/encoding/fuzz"
	"github.com/alxarno/parquet-go-athena-list-hack/encoding/rle"
)

func FuzzEncodeLevels(f *testing.F) {
	fuzz.EncodeLevels(f, &rle.Encoding{BitWidth: 8})
}
