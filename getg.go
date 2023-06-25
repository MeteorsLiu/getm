package getm

import (
	"unsafe"

	_ "github.com/v2pro/plz/gls"
	"github.com/v2pro/plz/reflect2"
)

var mOffset uintptr
var mIDOffset uintptr

func init() {
	gType := reflect2.TypeByName("runtime.g").(reflect2.StructType)
	if gType == nil {
		panic("failed to get runtime.g type")
	}
	mOffset = gType.FieldByName("m").Offset()
	mType := reflect2.TypeByName("runtime.m").(reflect2.StructType)
	if gType == nil {
		panic("failed to get runtime.g type")
	}
	mIDOffset = mType.FieldByName("id").Offset()
}

//go:linkname GetG github.com/v2pro/plz/gls.getg
func GetG() uintptr

func GetM() uintptr {
	g := GetG()
	m := (*uintptr)(unsafe.Pointer(g + mOffset))
	return *m
}

func MID() int64 {
	m := GetM()
	id := (*int64)(unsafe.Pointer(m + mIDOffset))
	return *id
}
