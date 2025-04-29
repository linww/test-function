package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aobco/log"
	uuid2 "github.com/google/uuid"
	"github.com/olivere/elastic/v7"
	"go.uber.org/atomic"
	"math/rand"
	"strings"
	"sync"
	"time"
	"unsafe"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/olivere"
	"xsky.com/ocpf/pkg/metadata/util"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	producerNum   = 200
	consumerNum   = 20
)

var (
	mtimeUtcDatePool      = make([]string, 4320)
	rcmtimeDatePool       = make([]string, 4320)
	rcmtimeUtcDatePool    = make([]string, 4320)
	lineNoPool            = make([]string, 20)
	machineNoPool         = make([]string, 100)
	cameraNoPool          = make([]string, 500)
	agentPool             = make([]string, 100)
	totalCount            = atomic.NewInt64(0)
	producerFinishedCount = atomic.NewInt64(0)
	targetCount           = int64(0)
	upsertBatch           = 1000
	standardKeyLength     = true
)

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

var (
	esUrl     string
	indexName string
)

func init() {
	flag.StringVar(&esUrl, "es_url", "", "es url")
	flag.StringVar(&indexName, "index_name", "", "index name")
	flag.Int64Var(&targetCount, "target_count", 0, "target count")
	flag.BoolVar(&standardKeyLength, "standard_key_length", true, "standard key length")
	flag.Parse()
	//esUrl = "http://127.0.0.1:9200"
	//indexName = "ocpf-eos-metadata-1-9999-127"
	//targetCount = 10000
	standardKeyLength = false
	if esUrl == "" {
		log.Fatalf("es ip is empty")
	}
	if indexName == "" {
		log.Fatalf("index name is empty")
	}
	if targetCount == 0 {
		log.Fatalf("target count is 0")
	}
	log.Infof("es url: %s, index name: %s, target count: %d, standard key length: %t ", esUrl, indexName, targetCount, standardKeyLength)
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	halfYearStart := now.AddDate(0, 0, 7).AddDate(0, 0, -180)
	halfYearStartHour := time.Date(halfYearStart.Year(), halfYearStart.Month(), halfYearStart.Day(), 0, 0, 0, 0, time.Local)
	for i := 0; i < 4320; i++ {
		date := halfYearStartHour.Add(time.Duration(i+1) * time.Hour).Add(time.Duration(rand.Int63n(60)) * time.Minute)
		utcDate := date.UTC().Format("2006-01-02T15:04:05.000000Z")
		mtimeUtcDatePool[i] = utcDate
	}
	for i := 0; i < 4320; i++ {
		date := halfYearStartHour.Add(time.Duration(i+1) * time.Hour).Add(time.Duration(rand.Int63n(60)) * time.Minute)
		rcmtimeDatePool[i] = date.Format("20060102150405")
		utcDate := date.UTC().Format("2006-01-02T15:04:05.000000Z")
		rcmtimeUtcDatePool[i] = utcDate
	}
	for i := 0; i < 20; i++ {
		if standardKeyLength {
			lineNoPool[i] = fmt.Sprintf("lineNo_%s", RandStringBytesMaskImprSrcUnsafe(57))
		} else {
			lineNoPool[i] = fmt.Sprintf("lineNo_%s", RandStringBytesMaskImprSrcUnsafe(25))
		}
	}
	for i := 0; i < 100; i++ {
		if standardKeyLength {
			machineNoPool[i] = fmt.Sprintf("machineNo_%s", RandStringBytesMaskImprSrcUnsafe(54))
		} else {
			machineNoPool[i] = fmt.Sprintf("machineNo_%s", RandStringBytesMaskImprSrcUnsafe(22))
		}
	}
	for i := 0; i < 500; i++ {
		if standardKeyLength {
			cameraNoPool[i] = fmt.Sprintf("cameraNo_%s", RandStringBytesMaskImprSrcUnsafe(55))
		} else {
			cameraNoPool[i] = fmt.Sprintf("cameraNo_%s", RandStringBytesMaskImprSrcUnsafe(23))
		}
	}
	for i := 0; i < 100; i++ {
		agentPool[i] = fmt.Sprintf("%s@%s", RandStringBytesMaskImprSrcUnsafe(1), uuid2.Must(uuid2.NewRandom()).String())
	}
}

func main() {
	esIpPool := getEsIpPool()
	wg := &sync.WaitGroup{}
	docChan := make(chan []elastic.BulkableRequest, 100)
	for i := 0; i < producerNum; i++ {
		wg.Add(1)
		go produceObjectDoc(wg, docChan)
	}
	for i := 0; i < consumerNum; i++ {
		wg.Add(1)
		esIp := esIpPool[i%len(esIpPool)]
		esClient, err := olivere.GetOrCreateClient(&es.Config{
			Url:      []string{esIp},
			Username: "elastic",
			Password: "123456",
		})
		if err != nil {
			log.Fatalf("get es client failed, %s", err)
		}
		go consumeObjectDoc(wg, docChan, esClient, esIp, i)
	}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			log.Infof("chan length: %d", len(docChan))
		}
	}()
	wg.Wait()
	log.Infof("total count: %d", totalCount.Load())
}

func consumeObjectDoc(wg *sync.WaitGroup, docChan chan []elastic.BulkableRequest, client *elastic.Client, esIp string, consumerNo int) {
	defer wg.Done()
	for {
		select {
		case requests, ok := <-docChan:
			if !ok {
				log.Infof("consume finished")
				return
			}
			bulkRequest := client.Bulk().Index(indexName)
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
			log.Infof("consumerNo: [%d], url: [%s] bulk request size [%d] success, cost: [%dms]", consumerNo, esIp, requestLen, time.Since(now).Milliseconds())
		}
	}
}

func produceObjectDoc(wg *sync.WaitGroup, docChan chan []elastic.BulkableRequest) {
	defer func() {
		wg.Done()
		if finishedCount := producerFinishedCount.Inc(); finishedCount == producerNum {
			close(docChan)
		}
	}()
	requests := make([]elastic.BulkableRequest, 0)
	for {
		if total := totalCount.Inc(); total > targetCount {
			break
		}
		mtimeRand := rand.Intn(4320)
		var rcMtimeRand int
		if mtimeRand == 0 {
			rcMtimeRand = 0
		} else {
			rcMtimeRand = rand.Intn(mtimeRand)
		}
		var keyRandStr string
		if standardKeyLength {
			keyRandStr = RandStringBytesMaskImprSrcUnsafe(106)
		} else {
			keyRandStr = RandStringBytesMaskImprSrcUnsafe(42)
		}
		key := fmt.Sprintf("%s/object_%s", rcmtimeDatePool[rcMtimeRand], keyRandStr)
		fileName := fmt.Sprintf("object_%s", keyRandStr)
		objectVersion := "null"
		//var tSn, mM1, mM2, mM3 string
		var tSn string
		if standardKeyLength {
			tSn = fmt.Sprintf("sn_%s", RandStringBytesMaskImprSrcUnsafe(105))
		} else {
			tSn = fmt.Sprintf("sn_%s", RandStringBytesMaskImprSrcUnsafe(51))
		}
		agentRand := rand.Intn(len(agentPool))
		doc := map[string]interface{}{
			"t_lineNo":       lineNoPool[rand.Intn(20)],
			"t_machineNo":    machineNoPool[rand.Intn(100)],
			"t_camera":       cameraNoPool[rand.Intn(500)],
			"t_sn":           tSn,
			"t_rc":           agentPool[agentRand],
			"m_rcmtime":      rcmtimeUtcDatePool[rcMtimeRand],
			"file_name":      fileName,
			"object_version": "null",
			"key":            key,
			"mtime":          mtimeUtcDatePool[mtimeRand],
			"size":           rand.Intn(101),
		}
		id := util.GenerateMetadataHashId(key, objectVersion)
		request := elastic.NewBulkIndexRequest().Id(id).Doc(doc)
		requests = append(requests, request)
		if len(requests) == upsertBatch {
			docChan <- requests
			requests = make([]elastic.BulkableRequest, 0)
		}
	}
	if len(requests) > 0 {
		docChan <- requests
	}
}

func getEsIpPool() []string {
	esIpPool := make([]string, 0)
	esConfig := &es.Config{
		Url:      []string{esUrl},
		Username: "elastic",
		Password: "123456",
	}
	client, err := olivere.GetOrCreateClient(esConfig)
	if err != nil {
		log.Fatalf("get es client failed, %s", err)
	}
	resp, err := client.NodesStats().Do(context.TODO())
	if err != nil {
		log.Fatalf("get es nodes stats failed, %s", err)
	}
	for _, node := range resp.Nodes {
		esIpPool = append(esIpPool, fmt.Sprintf("http://%s:9200", strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(node.IP, ":9300"), ":9301"), ":9302")))
	}
	esIpPool = append(esIpPool, esUrl)

	return esIpPool
}
