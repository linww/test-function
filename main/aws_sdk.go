package main

import (
	"github.com/aws/aws-sdk-go/aws"
	awscred "github.com/aws/aws-sdk-go/aws/credentials"
	awsrequest "github.com/aws/aws-sdk-go/aws/request"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	cred := awscred.NewStaticCredentials("11111111111111111111", "1111111111111111111111111111111111111111", "")
	cfg := aws.NewConfig().WithRegion("us-east-1").WithEndpoint("http://10.16.12.135:8060").
		WithS3ForcePathStyle(true).WithDisableSSL(true).
		WithCredentials(cred).
		WithMaxRetries(3)
	session, err := awssession.NewSession(cfg)
	if err != nil {
		panic(err)
	}
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(session)
	f, err := os.Create("/Users/xsky/Desktop/working/ocpf/main_test/main/example_resize_oss.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	WithHeader := func(key, val string) awsrequest.Option {
		return func(req *awsrequest.Request) {
			req.HTTPRequest.Header.Set(key, val)
		}
	}
	// Write the contents of S3 Object to the file
	command := "image/resize,m_fixed,h_100,w_100"
	option := s3manager.WithDownloaderRequestOptions(
		WithHeader("x-oss-process", command),
		WithHeader("x-amz-meta-datasource", "100000"))
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("ww-bucket"),
		Key:    aws.String("image.jpeg"),
	}, option)
	if err != nil {
		panic(err)
	}
}
