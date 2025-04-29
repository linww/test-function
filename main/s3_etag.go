package main

import (
	"context"
	"fmt"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	types2 "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"os"
	"sync"
	"xsky.com/ocpf/pkg/osg/service"
)

func main() {
	client := service.OsgRequestService.GetS3Client("gh1c84wtnqbxalwzonlf", "a9rttrsj1mu0r57op6v1yewszldfa2vsdynvz3x3", "http://10.16.17.55:8060")
	var nextToken string
	md5File, err := os.OpenFile("md5_files.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := md5File.Close(); err != nil {
			log.Errorf("close file %s error: %v", "md5_files.txt", err)
		}
	}()
	totalCount := 0
	objectChan := make(chan types2.Object, 10000)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(waitGroup *sync.WaitGroup, objectChan chan types2.Object) {
		defer waitGroup.Done()
		for {
			select {
			case obj, ok := <-objectChan:
				if !ok {
					return
				}
				_, err = md5File.WriteString(fmt.Sprintf("%s,%s\n", *obj.Key, *obj.ETag))
			}
		}

	}(wg, objectChan)
	for {
		output, err := client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
			Bucket:            aws.String("bucket-4"),
			ContinuationToken: aws.String(nextToken),
			MaxKeys:           10000,
		})
		if err != nil {
			log.Errorf("list objects error: %v", err)
			return
		}
		totalCount += len(output.Contents)
		for _, obj := range output.Contents {
			objectChan <- obj
		}
		log.Infof("total count: %d", totalCount)
		if !output.IsTruncated {
			break
		}
		nextToken = *output.NextContinuationToken
	}
	close(objectChan)
	wg.Wait()
	log.Infof("write md5 to file success")
}
