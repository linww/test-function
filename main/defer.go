package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		println(i)
		defer func() {
			println(fmt.Sprintf("defer %d", i))
		}()
	}
}
