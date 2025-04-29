package main

import (
	"fmt"
	"xsky.com/ocpf/pkg/license/model"
)

func main() {
	ll := &model.LicenseLimit{
		Id:        99999,
		Key:       "max_lto_capacity",
		Limits:    10995116277760,
		TotalUsed: 0,
		Hash:      "",
	}
	h, err := ll.CalculateHash()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)
}
