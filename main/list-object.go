package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	service2 "xsky.com/ocpf/pkg/osg/service"
)

func main() {
	var nextToken *string
	totalObjects := 0
	s3Client := service2.OsgRequestService.GetS3Client("11111111111111111111", "1111111111111111111111111111111111111111", "http://10.16.13.137:8060")
	for {
		output, err := s3Client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
			Bucket:            aws.String("bucket-3"),
			MaxKeys:           10000,
			ContinuationToken: nextToken,
		})
		if err != nil {
			panic(err)
		}
		totalObjects += len(output.Contents)
		if !output.IsTruncated {
			break
		}
		nextToken = output.NextContinuationToken
	}
	println(fmt.Sprintf("total objects: %d", totalObjects))
}
