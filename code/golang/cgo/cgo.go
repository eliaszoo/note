package main

/*
#include <stdio.h>
#include <stdlib.h>
void Print(char* s) {
	printf("%s\n", s);
}
*/
import "C"
import (
	"unsafe"
)

func main() {
	str := "Hello World"
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.Print(cstr)
}
