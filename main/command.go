package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "/Users/xsky/Desktop/working/ocpf/main_test/main/command.go")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	s := string(stdout)
	fmt.Println(s)
}
