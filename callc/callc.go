package callc

import (
	"fmt"
)

// #include <stdio.h>
//void callC() {
//    printf("call C code from golang\n");
//}
import "C"

func CallC() {
	fmt.Println("going to call c")
	C.callC()
	fmt.Println("call c done")
}


