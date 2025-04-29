package main

import "os"

func main() {
	if err := os.Rename("/Users/xsky/Downloads/xxx/a", "/Users/xsky/Downloads/xxx/b"); err != nil {
		panic(err)
	}
}
