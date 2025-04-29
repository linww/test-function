package main

import (
	"context"
	"github.com/aobco/log"
	uuid2 "github.com/google/uuid"
	"github.com/olivere/elastic/v7"
	"math/rand"
	"strings"
	"time"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/olivere"
)

func main() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).UTC()
	requests := make([]elastic.BulkableRequest, 0)
	esClient, err := olivere.GetOrCreateClient(&es.Config{
		Url:      []string{"http://localhost:9200"},
		Username: "elastic",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("get es client failed, %s", err)
	}
	//uuids := []string{"a310a25e-2f99-11ef-8a28-1e80201af071", "a310a286-2f99-11ef-8a2c-1e80201af071", "a310a27c-2f99-11ef-8a29-1e80201af071", "a310a272-2f99-11ef-8a28-1e80201af071", "a310a286-2f99-11ef-8a2b-1e80201af071", "a310a27c-2f99-11ef-8a2a-1e80201af071", "a310a290-2f99-11ef-8a2d-1e80201af071", "a310a27c-2f99-11ef-8a2b-1e80201af071", "a310a286-2f99-11ef-8a2d-1e80201af071", "a310a272-2f99-11ef-8a29-1e80201af071"}
	for i := 0; i < 20000; i++ {
		rcmtime := today.AddDate(0, 0, -rand.Intn(10))
		rcmtime = rcmtime.Add(time.Duration(-rand.Int63n(20)) * time.Hour)
		doc := map[string]interface{}{
			"key":       uuid2.Must(uuid2.NewUUID()).String(),
			"mtime":     today.AddDate(0, 0, -1).Add(time.Duration(rand.Int63n(20)) * time.Hour).Format("2006-01-02T15:04:05.000000Z"),
			"size":      10,
			"m_rcmtime": rcmtime.Format("2006-01-02T15:04:05.000000Z"),
			//"t_rc":      uuids[rand.Intn(10)],
		}
		request := elastic.NewBulkIndexRequest().Doc(doc)
		requests = append(requests, request)
		if len(requests) == 10000 {
			bulkRequest := esClient.Bulk().Index("ocpf-gfs-metadata-1-9998-9999")
			bulkRequest.Add(requests...)
			requestLen := bulkRequest.NumberOfActions()
			now := time.Now()
			resp, err := bulkRequest.Do(context.Background())
			if err != nil {
				if strings.Contains(err.Error(), "Request Entity Too Large") {
					log.Errorf("bulk request failed, request length: [%d], %v", requestLen, err)
					continue
				}
				log.Fatalf("bulk request failed, request length: [%d], %v", requestLen, err)
			}
			if resp.Errors {
				for _, item := range resp.Failed() {
					log.Errorf("bulk request failed, %s", item.Error.Reason)
				}
			}
			log.Infof("bulk request size [%d] success, cost: [%dms]", requestLen, time.Since(now).Milliseconds())
			requests = make([]elastic.BulkableRequest, 0)
		}
	}
}
