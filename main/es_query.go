package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aobco/log"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/olivere"
	response2 "xsky.com/ocpf/pkg/rest/response"
)

var (
	esUrl               string
	ocpfUrl             string
	esUsername          string
	esPassword          string
	queryStartTimestamp int64
	queryEndTimestamp   int64
	authToken           string
)

func init() {
	flag.StringVar(&esUrl, "es_url", "", "es url")
	flag.StringVar(&ocpfUrl, "ocpf_url", "", "ocpf url")
	flag.StringVar(&esUsername, "es_username", "", "es username")
	flag.StringVar(&esPassword, "es_password", "", "es password")
	flag.Int64Var(&queryStartTimestamp, "start_time", 0, "start time")
	flag.Int64Var(&queryEndTimestamp, "end_time", 0, "end time")
	flag.StringVar(&authToken, "auth_token", "", "auth token")
	flag.Parse()
}

type EsQueryResponse struct {
	response2.JsonResponse
	Content elastic.SearchResult `json:"content"`
}

var bodyTemplate = `
	{
    "index_name": "%s",
    "start_time": %d,
    "end_time": %d,
    "searchafter_key": "%s",
    "searchafter_time":%d
}
`

func main() {
	esIps := getEsIpPool()
	esConfig := &es.Config{
		Url:      esIps,
		Username: esUsername,
		Password: esPassword,
	}
	olivereClient, err := olivere.GetOrCreateClient(esConfig)
	if err != nil {
		log.Fatalf("get es client failed, %s", err)
	}
	catIndicesResp, err := olivereClient.CatIndices().Do(context.Background())
	if err != nil {
		log.Fatalf("get es cat indices failed, %s", err)
	}
	if err != nil {
		log.Fatalf("get es client failed, %s", err)
	}
	httpClient := http.DefaultClient
	allNow := time.Now()
	totalCount := 0
	for _, catIndex := range catIndicesResp {
		if !strings.HasPrefix(catIndex.Index, "ocpf-") {
			continue
		}
		body := fmt.Sprintf(bodyTemplate, catIndex.Index, queryStartTimestamp, queryEndTimestamp, "", 0)
		for {
			queryStart := time.Now()
			consumeUrl := fmt.Sprintf("%s/datasource/token/es/query-by-mtime", ocpfUrl)
			req, err := http.NewRequest(http.MethodPost, consumeUrl, strings.NewReader(body))
			if err != nil {
				log.Fatalf("create request failed, %s", err)
			}
			req.Header.Add("Xocp-Auth-Token", authToken)
			sendStart := time.Now()
			response, err := httpClient.Do(req)
			if err != nil {
				log.Fatalf("consume error %v", err)
			}
			sendUse := time.Since(sendStart).Milliseconds()
			readAllStart := time.Now()
			respBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatalf("read response body error %v", err)
			}
			readAllUse := time.Since(readAllStart).Milliseconds()
			response.Body.Close()
			if response.StatusCode != http.StatusOK {
				log.Fatalf("consume error %v", err)
			}
			unmarshalStart := time.Now()
			rep := &EsQueryResponse{}
			err = json.Unmarshal(respBytes, rep)
			if err != nil {
				log.Fatalf("unmarshal response error %v", err)
			}
			unmarshalUse := time.Since(unmarshalStart).Milliseconds()
			if len(rep.Content.Hits.Hits) == 0 {
				break
			}
			totalCount += len(rep.Content.Hits.Hits)
			requestId := response.Header.Get("Xocp-Request-Id")
			lastHit := rep.Content.Hits.Hits[len(rep.Content.Hits.Hits)-1]
			body = fmt.Sprintf(bodyTemplate, catIndex.Index, queryStartTimestamp, queryEndTimestamp, lastHit.Sort[1], int64(lastHit.Sort[0].(float64)))
			log.Infof("req: %s, consume %s, send %dms, readAll %dms, unmarshal %dms, total %dms", requestId, catIndex.Index, sendUse, readAllUse, unmarshalUse, time.Since(queryStart).Milliseconds())
		}
	}
	log.Infof("consume total count: %d use %dms", totalCount, time.Since(allNow).Milliseconds())

}

func getEsIpPool() []string {
	esIpPool := make([]string, 0)
	esConfig := &es.Config{
		Url:      []string{esUrl},
		Username: esUsername,
		Password: esPassword,
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

	return esIpPool
}
