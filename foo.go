package main

import "fmt"

func Foo(n int) int {
	var o int
	for i := 0; i < 10; i++ {
		o = o + i
	}
	return o
}

func main() { fmt.Println(Foo(1)) }
