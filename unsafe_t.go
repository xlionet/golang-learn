package main

import (
	"fmt"
	"unsafe"
)

func main() {
	u := uint32(1)
	i := int32(2)

	fmt.Println(&u, &i)

	p := &i
	p = (*int32)(&u)
	p = (*int32)(unsafe.Pointer(&u))

	fmt.Println(p)
}
