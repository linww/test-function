package main

import (
	"context"
	"fmt"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/atomic"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"xsky.com/ocpf/pkg/osg/service"
)

func main() {
	//s3Client := service.OsgRequestService.GetS3Client("gh1c84wtnqbxalwzonlf", "a9rttrsj1mu0r57op6v1yewszldfa2vsdynvz3x3", "http://10.16.17.55:8060")
	//reader, err := os.OpenFile("output.png", os.O_RDONLY, 0666)
	//if err != nil {
	//	log.Errorf("%v", err)
	//	return
	//}
	//_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
	//	Bucket: aws.String("ww-bucket"),
	//	Key:    aws.String("output.png"),
	//	Body:   reader,
	//})
	//if err != nil {
	//	log.Errorf("%v", err)
	//}
	//log.Infof("put object successfully")

	//client := service.OsgRequestService.GetS3Client("99999999999999999999", "9999999999999999999999999999999999999999", "http://10.16.13.137:8060")
	client := service.OsgRequestService.GetS3Client("11111111111111111111", "1111111111111111111111111111111111111111", "http://10.16.12.135:8060")
	putTotal := atomic.NewInt64(0)
	wg := &sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				current := putTotal.Inc()
				if current > 50000 {
					break
				}
				//current = current + 10000000
				rcmtime := time.Now().UnixMilli() - (rand.Int63n(2*24)+2)*time.Hour.Milliseconds()
				tags := make([]string, 0)
				tags = append(tags, fmt.Sprintf("rc=agent_%d", current))
				tags = append(tags, fmt.Sprintf("long=%d", current))
				tags = append(tags, fmt.Sprintf("double=%s", fmt.Sprintf("%.2f", float64(current)*0.1)))
				tags = append(tags, fmt.Sprintf("date=%s", time.UnixMilli(rcmtime).Format("20060102150405")))
				_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
					Bucket: aws.String("ww-bucket"),
					Key:    aws.String(fmt.Sprintf("archive/file_%d.txt", current+100000)),
					Body:   strings.NewReader(fmt.Sprintf("hello %d", current)),
					Metadata: map[string]string{
						"rcmtime": strconv.FormatInt(rcmtime, 10),
						"keyword": fmt.Sprintf("m_keyword_%d", current),
						"long":    strconv.FormatInt(current, 10),
						"double":  fmt.Sprintf("%.2f", float64(current)*0.1),
					},
					Tagging: aws.String(strings.Join(tags, "&")),
				})
				if err != nil {
					log.Errorf("put object failed: %v", err)
				}
			}
		}()
	}
	wg.Wait()
}
