package callc

// #cgo CFLAGS: -I${SRCDIR}/callc
// #cgo LDFLAGS: ${SRCDIR}/libcallC.a
// #include "callC.h"
// #include <stdlib.h>
import "C"
import "unsafe"

func CHello(){
	C.cHello()
}

func PrintMessage(message string){
	msg := C.CString(message)
	C.printMessage(msg)
	defer C.free(unsafe.Pointer(msg))
}
