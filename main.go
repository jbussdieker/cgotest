package main
/*
#include <stdio.h>
#include <stdlib.h>

void myprintf(char *s) {
	printf("%s", s);
}
*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("Test\n")
	C.myprintf(cs)
	println("Test")
	C.free(unsafe.Pointer(cs))
}
