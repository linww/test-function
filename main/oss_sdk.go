package main

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	client, err := oss.New("10.16.12.135:8060", "99999999999999999999", "9999999999999999999999999999999999999999")
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket("xsky-1")
	if err != nil {
		panic(err)
	}
	process := fmt.Sprintf("image/resize,w_100|sys/saveas,o_aW1hZ2Vfc2luay5qcGVn")
	_, err = bucket.ProcessObject("sea.jpeg", process, oss.AddParam("x-amz-meta-datasource", "100002"))
	if err != nil {
		panic(err)
	}
}
