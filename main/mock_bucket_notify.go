package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"time"
	osgService "xsky.com/ocpf/pkg/osg/service"
)

type BucketNotify struct {
	Records []*Record `json:"Records"`
}

type Record struct {
	UserIdentity *UserIdentity `json:"userIdentity"`
	S3           *S3           `json:"s3"`
}

type UserIdentity struct {
	PrincipalId string `json:"principalId"`
}

type S3 struct {
	Bucket *Bucket `json:"bucket"`
	Object *Object `json:"object"`
}

type Bucket struct {
	Name string `json:"name"`
}

type Object struct {
	Key  string `json:"key"`
	Size int    `json:"size"`
	Etag string `json:"eTag"`
}

var (
	srcBucket      string
	srcPrefix      string
	notifyInterval int64
)

func main() {
	flag.StringVar(&srcBucket, "src_bucket", "", "source bucket")
	flag.StringVar(&srcPrefix, "src_prefix", "", "source prefix")
	flag.Int64Var(&notifyInterval, "notify_interval", 5, "notify interval")
	flag.Parse()
	var objects []types.Object
	s3Client := osgService.OsgRequestService.GetS3Client("99999999999999999999", "9999999999999999999999999999999999999999", "http://10.16.12.135:8060")
	nextToken := ""
	for {
		listObjectsV2Output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket:            &srcBucket,
			Prefix:            &srcPrefix,
			ContinuationToken: &nextToken,
		})
		if err != nil {
			panic(err)
		}
		for _, content := range listObjectsV2Output.Contents {
			objects = append(objects, content)
		}
		if !listObjectsV2Output.IsTruncated {
			break
		}
		nextToken = *listObjectsV2Output.NextContinuationToken
	}
	log.Infof("src_bucket: %s, src_prefix, %s, notify_interval: %d, objects length: %d", srcBucket, srcPrefix, notifyInterval, len(objects))
	for {
		bucketNotify := &BucketNotify{}
		records := make([]*Record, 0)
		for i := 0; i < 999; i++ {
			object := objects[rand.Int63n(int64(len(objects)))]
			records = append(records, &Record{
				UserIdentity: &UserIdentity{
					PrincipalId: "user",
				},
				S3: &S3{
					Bucket: &Bucket{
						Name: srcBucket,
					},
					Object: &Object{
						Key:  *object.Key,
						Size: int(object.Size),
						Etag: *object.ETag,
					},
				},
			})
		}
		bucketNotify.Records = records
		if err := sendBucketNotify(bucketNotify); err != nil {
			log.Errorf("send bucket notify failed, %v", err)
		}
		time.Sleep(time.Duration(notifyInterval) * time.Second)
	}
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
		DialContext: (&net.Dialer{}).DialContext,
	},
	Timeout: 0,
}

func sendBucketNotify(bucketNotify *BucketNotify) error {
	targetUrl := "http://10.16.64.16/workflow/anon/notify/arn:aws:lambda:default:1:function:image-encoder?datasource_id=100120&rule_name=jpg"
	log.Debugf(targetUrl)
	body, err := json.Marshal(bucketNotify)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, targetUrl, bytes.NewReader(body))
	if err != nil {
		return err
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New(string(respBytes))
	}
	log.Infof("send bucket notify successfully")
	return nil
}
