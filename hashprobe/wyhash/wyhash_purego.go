//go:build purego || !amd64

package wyhash

import "github.com/alxarno/parquet-go-athena-list-hack/sparse"

func MultiHashUint32Array(hashes []uintptr, values sparse.Uint32Array, seed uintptr) {
	for i := range hashes {
		hashes[i] = Hash32(values.Index(i), seed)
	}
}

func MultiHashUint64Array(hashes []uintptr, values sparse.Uint64Array, seed uintptr) {
	for i := range hashes {
		hashes[i] = Hash64(values.Index(i), seed)
	}
}

func MultiHashUint128Array(hashes []uintptr, values sparse.Uint128Array, seed uintptr) {
	for i := range hashes {
		hashes[i] = Hash128(values.Index(i), seed)
	}
}
