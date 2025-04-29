package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string)
	wg := &sync.WaitGroup{}
	s := make([]string, 0)
	wg.Add(1)
	go func(ma map[string]string, slice []string) {
		ma["a"] = "b"
		slice = append(slice, "a")
		wg.Done()
	}(m, s)
	wg.Wait()
	println(m)

	st := []string{"1", "2", "3"}
	// 转换为 JSON 格式的字符串
	jsonStr, err := json.Marshal(st)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonStr))
}
