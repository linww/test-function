package main

import (
	"context"
	"strings"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/olivere"
)

func main() {
	esClient, err := olivere.GetOrCreateClient(&es.Config{
		Url: strings.Split("http://10.16.53.185", ","),
	})
	if err != nil {
		panic(err)
	}
	result, err := esClient.Count("ocpf-*").Do(context.Background())
	if err != nil {
		panic(err)
	}
	println(result)
}
