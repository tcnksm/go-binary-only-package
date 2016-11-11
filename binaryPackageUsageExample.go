/*
You should put this usage file with fake minimal source code, which now is
src/github.com/tcnksm/hello/hello.go.
*/
package main

import (
	"fmt"
	"hello"
)

func main() {
	fmt.Println(" start program ")
	fmt.Println(" print program :", hello.Hello("hello world"))
}
