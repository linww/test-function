package main

import (
	"context"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"sync"
	"xsky.com/ocpf/pkg/osg/service"
)

func main() {
	bucket := "ww-bucket"
	prefix := "archive"
	//client := service.OsgRequestService.GetS3Client("11111111111111111111", "1111111111111111111111111111111111111111", "http://10.16.12.135:8060")
	client := service.OsgRequestService.GetS3Client("11111111111111111111", "1111111111111111111111111111111111111111", "http://10.16.12.135:8060")
	nextToken := ""
	waitGroup := &sync.WaitGroup{}
	objectChan := make(chan types.Object, 10000)
	for i := 0; i < 50; i++ {
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				object, ok := <-objectChan
				if !ok {
					break
				}
				_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
					Bucket: aws.String(bucket),
					Key:    object.Key,
				})
				if err != nil {
					panic(err)
				}
				log.Infof("delete object [%s] in bucket [%s] successfully", *object.Key, bucket)
			}
		}(waitGroup)
	}
	for {
		listObjects, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket:            aws.String(bucket),
			Prefix:            aws.String(prefix),
			ContinuationToken: aws.String(nextToken),
			MaxKeys:           10000,
		})
		if err != nil {
			panic(err)
		}
		for _, content := range listObjects.Contents {
			objectChan <- content
		}
		nextToken = aws.ToString(listObjects.NextContinuationToken)
		if !listObjects.IsTruncated {
			break
		}
		if nextToken == "" {
			break
		}
	}
	close(objectChan)
	waitGroup.Wait()
	log.Infof("delete all objects in bucket [%s] with prefix [%s]", bucket, prefix)
}
