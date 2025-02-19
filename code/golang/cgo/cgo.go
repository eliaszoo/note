package main

/*
#include <stdio.h>
#include <stdlib.h>
void Print(char* s);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export Print
func Print(s *C.char) {
	fmt.Println(C.GoString(s))
}

func main() {
	str := "Hello World"
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	by := C.GoBytes(unsafe.Pointer(cstr), C.int(len(str)))
	_ = by

	C.Print(cstr)

}
