// Package main provides ...
package main

import (
	"fmt"
	"unsafe"
)

type V struct {
	i int32
	j int64
}

func (v V) GetI() {
	fmt.Printf("i = %+v\n", v.i)
}

func (v V) GetJ() {
	fmt.Printf("j = %+v\n", v.j)
}
func main() {
	var v *V = &V{199, 299}
	var i *int32 = (*int32)(unsafe.Pointer(v))

	fmt.Printf("i = %+v\n", i)

	fmt.Printf("value of uintptr: = %+v\n", uintptr(unsafe.Pointer(i)))

	*i = int32(98)

	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0)))))
	*j = int64(963)

	v.GetI()
	v.GetJ()
	//TODO
}
