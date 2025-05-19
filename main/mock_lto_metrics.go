package main

import (
	"context"
	"github.com/aobco/log"
	"github.com/olivere/elastic/v7"
	"math/rand"
	"sync"
	"time"
	"xsky.com/ocpf/console/service/archive/metrics"
	constant2 "xsky.com/ocpf/pkg/constant"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/olivere"
)

func main() {
	totalCount := 0
	docChan := make(chan *metrics.LtoMetricsRecord, 100)
	requests := make([]elastic.BulkableRequest, 0)
	esClient, err := olivere.GetOrCreateClient(&es.Config{
		Url: []string{"http://10.16.64.117"},
	})
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			select {
			case record, ok := <-docChan:
				if !ok {
					if len(requests) > 0 {
						_, err = elastic.NewBulkService(esClient).Add(requests...).Do(context.TODO())
						if err != nil {
							panic(err)
						}
					}
					wg.Done()
					return
				}
				request := elastic.NewBulkIndexRequest().Index("ocpf-metrics-lto-1").Doc(record)
				requests = append(requests, request)
				if len(requests) >= 1000 {
					_, err = elastic.NewBulkService(esClient).Add(requests...).Do(context.TODO())
					if err != nil {
						panic(err)
					}
					requests = make([]elastic.BulkableRequest, 0)
				}
			}
		}
	}()
	for {
		totalCount += 1
		if totalCount > 123450 {
			break
		}
		rand.Seed(time.Now().UnixMilli())
		datasourceId := int64(100265)
		archivePlanId := int64(100008)
		indicatorValue := rand.Int63n(11) + 190
		subtractMin := rand.Int63n(7 * 24 * 60)
		record := &metrics.LtoMetricsRecord{
			DatasourceId:   datasourceId,
			ArchivePlanId:  archivePlanId,
			Time:           time.Now().Add(-time.Duration(subtractMin) * time.Minute).UTC().Format(constant2.EsDefaultGoDateFormat),
			IndicatorName:  metrics.IndicatorNameWriteSpeed,
			IndicatorValue: float64(indicatorValue),
		}
		docChan <- record
	}
	close(docChan)
	wg.Wait()
	log.Infof("totalCount: %d", totalCount)
}
