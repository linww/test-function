package main

import (
	"bytes"
	"context"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"golang.org/x/sync/semaphore"
	"os"
	"sync"
	"xsky.com/ocpf/pkg/osg/service"
)

func main() {
	// aws s3 multipart uploader
	s3Client := service.OsgRequestService.GetS3Client("99999999999999999999", "9999999999999999999999999999999999999999", "http://10.16.13.137:8060")
	//bucket := "ww-prod-bucket-072901"
	//bucket := "easydata-dst-archive-66-100020"
	//key := "100m"
	//key := "D/2024-07-26-wasqaf4i2Dfyj6oAqxfVfBbk4R1lZ9wW"
	var parts []types.CompletedPart
	//largeObject, err := os.Open("/Users/xsky/Downloads/tmp/600m")
	fc := func(options *s3.Options) {
		options.APIOptions = []func(*middleware.Stack) error{
			smithyhttp.AddHeaderValue("x-sds-archive-type", "archive-obj"),
		}
	}
	//largeObject, err := os.Open("/var/folders/2w/t7smdmm963g4jshwpls071tc0000gn/T/large-object2802981245")
	largeObject, err := os.Open("100G.mp4")
	fileStat, _ := largeObject.Stat()
	// upload part
	partSize := int64(50 * 1024 * 1024) // 50MB
	partCount := (fileStat.Size() + partSize - 1) / partSize
	multipartUpload, err := s3Client.CreateMultipartUpload(context.Background(), &s3.CreateMultipartUploadInput{
		Bucket: aws.String("ww-bucket"),
		Key:    aws.String("100G.mp4"),
	}, fc)
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	sem := semaphore.NewWeighted(10)
	l := &sync.Mutex{}
	for i := int64(1); i <= partCount; i++ {
		partStart := (i - 1) * partSize
		if fileStat.Size()-partStart < partSize {
			partSize = fileStat.Size() - partStart
		}
		partBody := make([]byte, partSize)
		_, err = largeObject.ReadAt(partBody, partStart)
		if err != nil {
			panic(err)
		}

		if err := sem.Acquire(context.Background(), 1); err != nil {
			panic(err)
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, sem *semaphore.Weighted, lock *sync.Mutex, partNumber int64, partBody []byte) {
			defer func() {
				wg.Done()
				sem.Release(1)
			}()
			partInput := &s3.UploadPartInput{
				Bucket:     aws.String("ww-bucket"),
				Key:        aws.String("100G.mp4"),
				PartNumber: int32(partNumber),
				UploadId:   multipartUpload.UploadId,
				Body:       bytes.NewReader(partBody),
			}

			// 上传分段
			out, err := s3Client.UploadPart(context.Background(), partInput, fc)
			if err != nil {
				panic(err)
			}
			lock.Lock()
			defer lock.Unlock()
			parts = append(parts, types.CompletedPart{
				ETag:       out.ETag,
				PartNumber: int32(partNumber),
			})
			log.Infof("part %d done", partNumber)
		}(wg, sem, l, i, partBody)
	}
	wg.Wait()
	// 完成上传
	_, err = s3Client.CompleteMultipartUpload(context.Background(), &s3.CompleteMultipartUploadInput{
		Bucket:              aws.String("ww-bucket"),
		Key:                 aws.String("100G.mp4"),
		UploadId:            multipartUpload.UploadId,
		ExpectedBucketOwner: nil,
		MultipartUpload:     &types.CompletedMultipartUpload{Parts: parts},
		RequestPayer:        "",
	}, fc)
	if err != nil {
		panic(err)
	}
	log.Infof("upload done")
}
