package main

import "C"
import "fmt"

//export PrintMessageFromgo
func PrintMessageFromgo(){
	fmt.Println("PrintMessage from GO!")
}

//export Multiply
func Multiply(a, b int) int {

	return a * b
}

/*
  the main function is required to build a shared library, but it will not be called when the library is used in C code.
*/
func main() {

}