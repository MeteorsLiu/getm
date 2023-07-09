package getm

import (
	"unsafe"

	_ "github.com/v2pro/plz/gls"
	"github.com/v2pro/plz/reflect2"
)

var (
	gType     reflect2.StructType
	mType     reflect2.StructType
	pType     reflect2.StructType
	mOffset   uintptr
	pOffset   uintptr
	mIDOffset uintptr
)

func init() {
	gType = reflect2.TypeByName("runtime.g").(reflect2.StructType)
	if gType == nil {
		panic("failed to get runtime.g type")
	}
	mOffset = gType.FieldByName("m").Offset()
	mType = reflect2.TypeByName("runtime.m").(reflect2.StructType)
	if mType == nil {
		panic("failed to get runtime.m type")
	}
	mIDOffset = mType.FieldByName("id").Offset()

	pType = reflect2.TypeByName("runtime.p").(reflect2.StructType)
	if pType == nil {
		panic("failed to get runtime.p type")
	}
	pOffset = mType.FieldByName("p").Offset()
}

//go:linkname GetG github.com/v2pro/plz/gls.getg
func GetG() uintptr

func GetM() uintptr {
	g := GetG()
	m := (*uintptr)(unsafe.Pointer(g + mOffset))
	return *m
}

func GetP() uintptr {
	m := GetM()
	p := (*uintptr)(unsafe.Pointer(m + pOffset))
	return *p
}

func MID() int64 {
	m := GetM()
	id := (*int64)(unsafe.Pointer(m + mIDOffset))
	return *id
}

// use at your own risks.
func CustomInG[T any](fieldName string) T {
	customOffset := gType.FieldByName(fieldName).Offset()
	return *(*T)(unsafe.Pointer(GetG() + customOffset))
}

// use at your own risks.
func CustomInM[T any](fieldName string) T {
	customOffset := mType.FieldByName(fieldName).Offset()
	return *(*T)(unsafe.Pointer(GetM() + customOffset))
}

// use at your own risks.
func CustomInP[T any](fieldName string) T {
	customOffset := pType.FieldByName(fieldName).Offset()
	return *(*T)(unsafe.Pointer(GetP() + customOffset))
}

// use at your own risks.
func SetCustomInG[T any](fieldName string, value T) {
	customOffset := gType.FieldByName(fieldName).Offset()
	*(*T)(unsafe.Pointer(GetG() + customOffset)) = value
}

// use at your own risks.
func SetCustomInM[T any](fieldName string, value T) {
	customOffset := mType.FieldByName(fieldName).Offset()
	*(*T)(unsafe.Pointer(GetM() + customOffset)) = value
}

// use at your own risks.
func SetCustomInP[T any](fieldName string, value T) {
	customOffset := pType.FieldByName(fieldName).Offset()
	*(*T)(unsafe.Pointer(GetP() + customOffset)) = value
}
