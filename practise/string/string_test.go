package practise

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {

	// 1. print string and length
	var str = "string value"

	fmt.Printf("string: %s, length: %d\n", str, len(str))

	// 2 iteral string
	for i, v := range str {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}

	// 3. string convert to bytes
	fmt.Printf("%s", "---------------------------------------------")
	var bytes []byte = []byte(str)

	for i, b := range bytes {
		fmt.Printf("index: %d, value: %d\n", i, b)
	}
	fmt.Printf("%s", "---------------------------------------------")
	// 4 string convert to rune

	var runes []rune = []rune(str)

	for i, b := range runes {
		fmt.Printf("index: %d, value: %d\n", i, b)
	}

	// 5. string join

	var arr []string = []string{"hello", "world", "from", "go"}

	fmt.Printf("%s", strings.Join(arr, "|"))

}
