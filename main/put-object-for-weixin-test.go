package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	osgService "xsky.com/ocpf/pkg/osg/service"
)

const (
	// key: GX_S_SMT_FAVI_0051/disk_I deepiresults/viresults/data/bside/20230410/13/134239_1696-HK_65M A/20230410134641585.HK_65M A134239_1696_0001r02c00.C47HIKBOVUF000000PK+GIVJ_OK_1.jpg
	keyTemplate        = "%s%s%s"
	fixedKeyPrefix     = "GX_S_SMT_FAVI_0051/disk_I deepiresults/viresults/data/bside/"
	fixedKeyComponents = "%s/%s/%s_1696-HK_65M A/"
	okFilename         = "%s.HK_65M A134239_1696_0001r02c00.C47HIKBOVUF000000PK+GIVJ_OK_%d.jpg"
	ngFilename         = "%s.HK_65M A134239_1696_0001r02c00.C47HIKBOVUF000000PK+GIVJ_NG_%d.jpg"
	weixinTimeFmt      = "20060102150405.999"
)

var (
	bucketName   string
	fileNamePool = []string{
		okFilename,
		ngFilename,
	}
	datePoolMap = map[string]int64{
		"20230410": 1000000,
		"20240410": 1000000,
		"20230411": 100000,
		"20240411": 100000,
		"20230412": 100000,
		"20240412": 100000,
		"20230413": 100000,
		"20240413": 100000,
		"20230414": 100000,
		"20240414": 100000,
		"20230415": 100000,
		"20240415": 100000,
		"20250414": 100000,
	}
	timePool = []string{
		"134239",
		"000000",
		"235959",
	}
	bucketTimesMap = map[string]float64{
		"case-1": 1,
		"case-2": 0.1,
		"case-3": 0.01,
		"case-4": 0.001,
		"case-5": 0.0001,
		"case-6": 0.00001,
	}
	weixinTotalCount = &atomic.Int64{}
	weixinS3Client   = &s3.Client{}
)

func main() {
	flag.StringVar(&bucketName, "bucket", "", "bucket name")
	flag.Parse()
	if bucketName == "" {
		log.Fatalf("bucket name is empty")
		return
	}
	weixinS3Client = osgService.OsgRequestService.GetS3Client("99999999999999999999", "9999999999999999999999999999999999999999", "http://10.16.12.14:8060")
	waitGroup := &sync.WaitGroup{}
	for date, count := range datePoolMap {
		waitGroup.Add(1)
		go putObjectByDate(waitGroup, date, count)
	}
	waitGroup.Wait()
	log.Infof("all done, total count: %d", weixinTotalCount.Load())

}

func putObjectByDate(waitGroup *sync.WaitGroup, fileDateStr string, count int64) {
	defer waitGroup.Done()
	bucketTimes := bucketTimesMap[bucketName]
	count = int64(float64(count) * bucketTimes)
	if count < 1 {
		log.Infof("bucket: %s, count: %d, bucketTimes: %f", bucketName, count, bucketTimes)
		return
	}
	innerWaitGroup := &sync.WaitGroup{}
	for _, fileTimeStr := range timePool {
		innerWaitGroup.Add(1)
		go putObjectByTime(innerWaitGroup, fileDateStr, fileTimeStr, count)
	}
	innerWaitGroup.Wait()
}

func getRandomMillisecond() string {
	rand.Seed(time.Now().Unix())
	mills := rand.Int63n(1000)
	if mills < 100 && mills > 9 {
		return fmt.Sprintf("0%d", mills)
	} else if mills < 10 {
		return fmt.Sprintf("00%d", mills)
	}
	return fmt.Sprintf("%d", mills)
}

func putObjectByTime(waitGroup *sync.WaitGroup, fileDateStr string, fileTimeStr string, count int64) {
	defer waitGroup.Done()
	innerWaitGroup := &sync.WaitGroup{}
	for _, filenameTpl := range fileNamePool {
		innerWaitGroup.Add(1)
		go putObjectByFileName(innerWaitGroup, filenameTpl, fileDateStr, fileTimeStr, count)
	}
	innerWaitGroup.Wait()
}

func putObjectByFileName(group *sync.WaitGroup, filenameTpl string, fileDateStr string, fileTimeStr string, count int64) {
	defer group.Done()
	sem := semaphore.NewWeighted(20)
	wg := &sync.WaitGroup{}
	for i := 0; i < int(count); i++ {
		currentFileNo := weixinTotalCount.Add(1)
		dateTimeStr := fmt.Sprintf("%s%s%s", fileDateStr, fileTimeStr, getRandomMillisecond())
		dateTimeFmt := dateTimeStr[:14] + "." + dateTimeStr[14:]
		fileTime, err := time.ParseInLocation(weixinTimeFmt, dateTimeFmt, time.Local)
		if err != nil {
			log.Fatalf("parse time error: %s", err)
			return
		}
		keyFixedComponents := fmt.Sprintf(fixedKeyComponents, fileDateStr, getFileTimeHour(fileTime), fileTimeStr)
		filename := fmt.Sprintf(filenameTpl, dateTimeStr, currentFileNo)
		key := fmt.Sprintf(keyTemplate, fixedKeyPrefix, keyFixedComponents, filename)
		log.Infof("key: %s", key)
		if err := sem.Acquire(context.TODO(), 1); err != nil {
			log.Fatalf("acquire semaphore error: %s", err)
			return
		}
		wg.Add(1)
		go func(sema *semaphore.Weighted, wg *sync.WaitGroup, bucketName string, key string, fileTime time.Time, currentFileNo int64) {
			defer wg.Done()
			defer sema.Release(1)
			_, err = weixinS3Client.PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket:   aws.String(bucketName),
				Key:      aws.String(key),
				Body:     strings.NewReader(fmt.Sprintf("%d", currentFileNo)),
				Metadata: map[string]string{"rcmtime": fmt.Sprintf("%d", fileTime.UnixMilli())},
			})
			if err != nil {
				log.Errorf("put object error: %s", err)
				return
			}
		}(sem, wg, bucketName, key, fileTime, currentFileNo)
	}
	wg.Wait()
}

func getFileTimeHour(fileTime time.Time) string {
	hour := fileTime.Hour()
	if hour < 10 {
		return fmt.Sprintf("0%d", hour)
	}
	return fmt.Sprintf("%d", hour)
}
