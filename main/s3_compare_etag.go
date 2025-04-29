package main

import (
	"bufio"
	"context"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	types2 "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"os"
	"strings"
	"sync"
	"xsky.com/ocpf/pkg/osg/service"
)

func main() {
	md5Map := make(map[string]string)
	md5File, err := os.OpenFile("md5_files.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Infof("open file error: %v", err)
		return
	}
	defer func() {
		if err := md5File.Close(); err != nil {
			log.Errorf("close file %s error: %v", "md5_files.txt", err)
		}
	}()
	lineReader := bufio.NewReader(md5File)
	for {
		line, err := lineReader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("read line error: %v", err)
			return
		}
		var key, md5 string
		lineSlice := strings.Split(strings.TrimSuffix(string(line), "\n"), ",")
		if len(lineSlice) != 2 {
			log.Errorf("line [%s] is invalid", string(line))
			return
		}
		key = lineSlice[0]
		md5 = lineSlice[1]
		md5Map[key] = md5
	}
	log.Infof("md5 map size: %d", len(md5Map))
	client := service.OsgRequestService.GetS3Client("gh1c84wtnqbxalwzonlf", "a9rttrsj1mu0r57op6v1yewszldfa2vsdynvz3x3", "http://10.16.17.55:8060")
	var nextToken string
	objectChan := make(chan types2.Object, 10000)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	prefix := "easydata/lto-archive/retrieve/100000/files/"
	equaledNum := 0
	notEqualNum := 0
	go func(waitGroup *sync.WaitGroup, objectChan chan types2.Object) {
		defer waitGroup.Done()
		for {
			select {
			case obj, ok := <-objectChan:
				if !ok {
					return
				}
				key := strings.TrimPrefix(*obj.Key, prefix)
				etag := md5Map[key]
				if etag != *obj.ETag {
					log.Warnf("object [%s] etag [%s] not equals to original etag [%s]", key, *obj.ETag, etag)
					notEqualNum++
				} else {
					equaledNum++
				}
			}
		}
	}(wg, objectChan)
	for {
		output, err := client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
			Bucket:            aws.String("bucket-1"),
			ContinuationToken: aws.String(nextToken),
			MaxKeys:           10000,
			Prefix:            aws.String("easydata/lto-archive/retrieve/100000"),
		})
		if err != nil {
			log.Errorf("list objects error: %v", err)
			return
		}
		for _, obj := range output.Contents {
			objectChan <- obj
		}
		if !output.IsTruncated {
			break
		}
		nextToken = *output.NextContinuationToken
	}
	close(objectChan)
	wg.Wait()
	log.Infof("compare finished, equaled num: %d, not equal num: %d", equaledNum, notEqualNum)
}
