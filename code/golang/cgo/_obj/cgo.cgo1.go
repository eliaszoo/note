// Code generated by cmd/cgo; DO NOT EDIT.

//line /Users/elias/docs/note/code/golang/cgo/cgo.go:1:1
package main

/*
#include <stdio.h>
#include <stdlib.h>
void Print(char* s);
*/
import _ "unsafe"
import (
	"fmt"
	"unsafe"
)

//export Print
func Print(s * /*line :15:15*/ _Ctype_char /*line :15:21*/) {
	fmt.Println(( /*line :16:14*/ _Cfunc_GoString /*line :16:23*/)(s))
}

func main() {
	str := "Hello World"
	cstr := ( /*line :21:10*/ _Cfunc_CString /*line :21:18*/)(str)
	defer func() func() {
		_cgo0 := /*line :22:15*/ unsafe.Pointer(cstr)
		return func() { _cgoCheckPointer(_cgo0, nil); /*line :22:36*/ _Cfunc_free(_cgo0) }
	}()()

	by := func() []byte {
		_cgo0 := /*line :24:18*/ unsafe.Pointer(cstr)
		var _cgo1 _Ctype_int = _Ctype_int /*line :24:45*/ (len(str))
		_cgoCheckPointer(_cgo0, nil)
		return /*line :24:56*/ _Cfunc_GoBytes(_cgo0, _cgo1)
	}()
	_ = by

	( /*line :27:2*/ _Cfunc_Print /*line :27:8*/)(cstr)
}
