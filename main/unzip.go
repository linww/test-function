package main

import (
	"archive/zip"
	"bytes"
	"context"
	"github.com/aobco/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"xsky.com/ocpf/pkg/osg/service"
)

func main() {
	client := service.OsgRequestService.GetS3Client("11111111111111111111", "1111111111111111111111111111111111111111", "http://10.16.12.135:8060")
	object, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("ww-bucket"),
		Key:    aws.String("periodic.zip"),
	})
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	defer func() {
		err := object.Body.Close()
		if err != nil {
			log.Errorf("%v", err)
			return
		}
	}()
	zipBytes, err := io.ReadAll(object.Body)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), object.ContentLength)
	if err != nil {
		log.Errorf("error when create zip reader %v", err)
		return
	}
	for _, file := range zipReader.File {
		fileReader, err := file.Open()
		if err != nil {
			log.Errorf("error when open file %v", err)
			return
		}
		defer fileReader.Close()
		fBytes, err := io.ReadAll(fileReader)
		if err != nil {
			log.Errorf("error when read file %v", err)
			return
		}
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("ww-bucket"),
			Key:    aws.String(file.Name),
			Body:   bytes.NewReader(fBytes),
		})
		if err != nil {
			log.Errorf("error when put object %v", err)
			return
		}
	}
}
