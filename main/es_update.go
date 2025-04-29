package main

import (
	context2 "context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aobco/log"
	"github.com/docker/distribution/context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/native"
	"xsky.com/ocpf/pkg/es/olivere"
)

var (
	queryDsl = `{
  "query": {
    "range": {
      "mtime": {
        "gte": "%s",
        "lt": "%s"
      }
    }
  }
}`
	querySort       = json.RawMessage(`[{"mtime":{"order":"desc"}}]`)
	querySize       = 10000
	queryEsUrl      = "http://es.ocpf.xsky.com"
	queryIndexName  = "ocpf-external-metadata-1-100002-100016"
	queryPrefix     = "/mnt/f20/"
	queryTotalCount = 0
	queryTotalCost  = time.Now()
	queryUsername   = ""
	queryPassword   = ""
	queryStartTime  = ""
	queryEndTime    = ""
)

func main() {
	flag.StringVar(&queryStartTime, "start_time", "", "start time")
	flag.StringVar(&queryEndTime, "end_time", "", "end time")
	flag.Parse()
	if queryStartTime == "" || queryEndTime == "" {
		log.Fatalf("invalid start time or end time")
	}
	iterator := new(DslIterator)
	queryDsl = fmt.Sprintf(queryDsl, queryStartTime, queryEndTime)
	err := iterator.SetUp([]string{queryIndexName}, queryDsl, querySort, querySize)
	if err != nil {
		log.Fatalf("setup error, %v", err)
	}
	esClient, err := olivere.CreateClient(&es.Config{
		Url:      []string{queryEsUrl},
		Username: queryUsername,
		Password: queryPassword,
	})
	requests := make([]elastic.BulkableRequest, 0)
	var id string
	for {
		item, err := iterator.Next()
		if err != nil {
			log.Fatalf("next error, %v", err)
		}
		if item == nil {
			break
		}
		if strings.HasPrefix(item["key"].(string), queryPrefix) {
			id = item["_id"].(string)
			key := strings.ReplaceAll(item["key"].(string), queryPrefix, "")
			vo := make(map[string]interface{}, 0)
			vo["key"] = key
			request := elastic.NewBulkUpdateRequest().Id(id).Doc(vo).DocAsUpsert(true)
			queryTotalCount++
			requests = append(requests, request)
		}
		if len(requests) == querySize {
			bulkRequest := esClient.Bulk().Index(queryIndexName)
			bulkRequest.Add(requests...)
			requestLen := bulkRequest.NumberOfActions()
			resp, err := bulkRequest.Do(context.Background())
			if err != nil {
				log.Fatalf("bulk request failed, request length: [%d], %v", requestLen, err)
			}
			if resp.Errors {
				for _, item := range resp.Failed() {
					log.Errorf("bulk request failed, %s", item.Error.Reason)
				}
			}
			log.Infof("id: [%s], total count: %d, total cost: %s", id, queryTotalCount, time.Since(queryTotalCost))
			requests = make([]elastic.BulkableRequest, 0)
		}
	}
	if len(requests) > 0 {
		bulkRequest := esClient.Bulk().Index(queryIndexName)
		bulkRequest.Add(requests...)
		requestLen := bulkRequest.NumberOfActions()
		resp, err := bulkRequest.Do(context.Background())
		if err != nil {
			log.Fatalf("bulk request failed, request length: [%d], %v", requestLen, err)
		}
		if resp.Errors {
			for _, item := range resp.Failed() {
				log.Errorf("bulk request failed, %s", item.Error.Reason)
			}
		}
		requests = make([]elastic.BulkableRequest, 0)
		log.Infof("total count: %d, total cost: %s", queryTotalCount, time.Since(queryTotalCost))
	}
}

type DslIterator struct {
	indices         []string
	dsl             string
	esClient        *elasticsearch.Client
	hits            []*elastic.SearchHit
	lastHit         *elastic.SearchHit
	currentPosition int
	sort            json.RawMessage
	size            int
}

func (d *DslIterator) SetUp(indices []string, dsl string, sort json.RawMessage, size int) error {
	d.indices = indices
	config := &es.Config{
		Url:      []string{queryEsUrl},
		Username: queryUsername,
		Password: queryPassword,
	}
	client, err := native.GetOrCreateClient(config)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	d.esClient = client
	dslQuery := make(map[string]interface{})
	if err = json.Unmarshal([]byte(dsl), &dslQuery); err != nil {
		log.Errorf("%v", err)
		return err
	}
	conditions := dslQuery["query"]
	conditionsDsl, err := json.Marshal(conditions)
	d.dsl = string(conditionsDsl)
	d.sort = sort
	d.size = size
	return nil
}

func (d *DslIterator) Next() (map[string]interface{}, error) {
	if len(d.hits) == 0 || d.currentPosition == len(d.hits) {
		getDsl, err := getDsl(d.dsl, d.lastHit, d.sort, d.size)
		if err != nil {
			log.Errorf("%v", err)
			return nil, err

		}
		search, err := d.esClient.Search( // Perform the search request.
			d.esClient.Search.WithContext(context2.Background()),
			d.esClient.Search.WithIndex(d.indices...),
			d.esClient.Search.WithBody(strings.NewReader(getDsl)),
			d.esClient.Search.WithTrackTotalHits(false),
		)
		if err != nil {
			log.Errorf("%v", err)
			return nil, err
		}
		if search.StatusCode != http.StatusOK {
			log.Warnf(search.String())
			return nil, fmt.Errorf("response unknown code:%d", search.StatusCode)
		}
		defer func() {
			err := search.Body.Close()
			if err != nil {
				log.Errorf("%v", err)
				return
			}
		}()
		all, err := ioutil.ReadAll(search.Body)
		if err != nil {
			log.Errorf("%v", err)
			return nil, err
		}
		result := &elastic.SearchResult{}
		err = json.Unmarshal(all, result)
		if err != nil {
			log.Errorf("%v", err)
			return nil, err
		}
		if len(result.Hits.Hits) == 0 {
			log.Infof("\n%s\n consume over", d.dsl)
			return nil, nil
		}
		d.hits = result.Hits.Hits
		d.currentPosition = 0
	}
	hit := d.hits[d.currentPosition]
	d.currentPosition++
	d.lastHit = hit

	item := make(map[string]interface{})
	err := json.Unmarshal(hit.Source, &item)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	item["index"] = hit.Index
	item["_id"] = hit.Id
	return item, nil
}

func (d *DslIterator) Close() error {
	return nil
}

type Dsl struct {
	Sort        json.RawMessage        `json:"sort"`
	Size        int                    `json:"size"`
	SearchAfter []interface{}          `json:"search_after,omitempty"`
	Query       map[string]interface{} `json:"query,omitempty"`
}

func getDsl(dsl string, hit *elastic.SearchHit, sort json.RawMessage, size int) (string, error) {
	d := new(Dsl)
	d.Sort = sort
	d.Size = size
	if hit != nil {
		d.SearchAfter = hit.Sort
	}
	if len(dsl) > 0 {
		query := make(map[string]interface{})
		err := json.Unmarshal([]byte(dsl), &query)
		if err != nil {
			log.Errorf("%v", err)
			return "", err
		}
		d.Query = query
	}
	indent, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Errorf("%v", err)
		return "", err
	}
	log.Infof("\n%s\n", string(indent))
	return string(indent), nil
}
