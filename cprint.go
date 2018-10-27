package main

import "C"
import "unsafe"

func main() {
	cstr := C.Cstring("Hello world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
