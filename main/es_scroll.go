package main

import (
	"encoding/json"
	"github.com/aobco/log"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/native"
)

func main() {
	client, err := native.GetOrCreateClient(&es.Config{
		Url:      []string{"http://10.16.53.185"},
		Username: "elastic",
		Password: "123456",
	})
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	search, err := client.Search(
		client.Search.WithIndex("ocpf-gfs-metadata-1-100002-100000,ocpf-eos-metadata-1-100003-3,ocpf-eos-metadata-1-100003-26,ocpf-eos-metadata-1-100003-20,ocpf-eos-metadata-1-100003-22,ocpf-external-metadata-1-100038-100140,ocpf-external-metadata-1-100038-100141,ocpf-external-metadata-1-100038-100142,ocpf-external-metadata-1-100038-100143,ocpf-external-metadata-1-100038-100144"),
		client.Search.WithBody(strings.NewReader(`{"query":{"match_all":{}},"sort":[{"mtime":{"order":"desc"}},{"key":{"order":"desc"}}]}`)),
		client.Search.WithSize(10000),
		client.Search.WithScroll(30*time.Minute),
	)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if search.StatusCode != http.StatusOK {
		log.Warnf(search.String())
		return
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
		return
	}
	result := &elastic.SearchResult{}
	err = json.Unmarshal(all, result)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	scrollId := result.ScrollId
	totalScrollCount := 0
	totalScrollCount += len(result.Hits.Hits)
	for {
		search, err := client.Scroll(
			client.Scroll.WithScrollID(scrollId),
			client.Scroll.WithScroll(30*time.Minute),
		)
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		if search.StatusCode != http.StatusOK {
			log.Warnf(search.String())
			return
		}
		//defer func() {
		//	err := search.Body.Close()
		//	if err != nil {
		//		log.Errorf("%v", err)
		//		return
		//	}
		//}()
		all, err := ioutil.ReadAll(search.Body)
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		result := &elastic.SearchResult{}
		err = json.Unmarshal(all, result)
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		if len(result.Hits.Hits) == 0 {
			break
		}
		totalScrollCount += len(result.Hits.Hits)
		scrollId = result.ScrollId
	}
	_, err = client.ClearScroll(
		client.ClearScroll.WithScrollID(scrollId),
	)
	if err != nil {
		log.Errorf("%v", err)
	}
	log.Infof("total scroll count: %d", totalScrollCount)
}
