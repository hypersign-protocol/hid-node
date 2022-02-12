package utils

import (
	"reflect"
	"strings"
	"unsafe"
)

// UnsafeStrToBytes uses unsafe to convert string into byte array. Returned bytes
// must not be altered after this function is called as it will cause a segmentation fault.
func UnsafeStrToBytes(s string) []byte {
	var buf []byte
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bufHdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	bufHdr.Data = sHdr.Data
	bufHdr.Cap = sHdr.Len
	bufHdr.Len = sHdr.Len
	return buf
}

// UnsafeBytesToStr is meant to make a zero allocation conversion
// from []byte -> string to speed up operations, it is not meant
// to be used generally, but for a specific pattern to delete keys
// from a map.
func UnsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ExtractIDFromSchema(id string) string {
	// Seperate `id` with seperator `;`
	extract_list := strings.Split(id, ";")

	// Second element of extract_list has the id.
	// Split that element with seperator `=`
	id_val := strings.Split(extract_list[1], "=")[1]
	return id_val
}
