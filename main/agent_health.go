package main

import (
	"context"
	"encoding/json"
	"github.com/aobco/log"
	"github.com/olivere/elastic/v7"
	"time"
	"xsky.com/ocpf/logservice/model"
	"xsky.com/ocpf/pkg/es"
	"xsky.com/ocpf/pkg/es/olivere"
)

func main() {
	agentJson := `{
		"health": {
	"cpu": 76.71232876712328,
	"mem": {
		"free": 15228928,
		"slab": 0,
		"used": 17164050432,
		"dirty": 0,
		"total": 17179279360,
		"wired": 0,
		"active": 0,
		"cached": 0,
		"mapped": 0,
		"shared": 0,
		"buffers": 0,
		"laundry": 0,
		"lowfree": 0,
		"highfree": 0,
		"inactive": 0,
		"lowtotal": 0,
		"swapfree": 0,
		"available": 15228928,
		"hightotal": 0,
		"swaptotal": 0,
		"writeback": 0,
		"pagetables": 0,
		"sunreclaim": 0,
		"swapcached": 0,
		"commitlimit": 0,
		"committedas": 0,
		"usedPercent": 99,
		"vmallocused": 0,
		"hugepagesize": 0,
		"sreclaimable": 0,
		"vmallocchunk": 0,
		"vmalloctotal": 0,
		"writebacktmp": 0,
		"hugepagesfree": 0,
		"hugepagestotal": 0
	},
	"disk": [{
		"free": 111891603456,
		"path": "C:",
		"used": 102180556800,
		"total": 214072160256,
		"fstype": "",
		"inodesFree": 0,
		"inodesUsed": 0,
		"inodesTotal": 0,
		"usedPercent": 47.731828687021476,
		"inodesUsedPercent": 0
	}, {
		"free": 241953300480,
		"path": "F:",
		"used": 80166096896,
		"total": 322119397376,
		"fstype": "",
		"inodesFree": 0,
		"inodesUsed": 0,
		"inodesTotal": 0,
		"usedPercent": 24.887075273652208,
		"inodesUsedPercent": 0
	}, {
		"free": 593389842432,
		"path": "G:",
		"used": 1605630263296,
		"total": 2199020105728,
		"fstype": "",
		"inodesFree": 0,
		"inodesUsed": 0,
		"inodesTotal": 0,
		"usedPercent": 73.01571545952035,
		"inodesUsedPercent": 0
	}, {
		"free": 5269241102336,
		"path": "H:",
		"used": 228298158080,
		"total": 5497539260416,
		"fstype": "",
		"inodesFree": 0,
		"inodesUsed": 0,
		"inodesTotal": 0,
		"usedPercent": 4.152733564338831,
		"inodesUsedPercent": 0
	}],
	"uuid": "036a7b72-b822-4409-ae0f-9d25f3ff5dd8",
	"gpu_info": null,
	"hostname": "DESKTOP-1BUS8QR",
	"networks": [{
		"ip": "10.16.111.1",
		"mac": "5a:4b:59:49:60:20"
	}, {
		"ip": "10.16.112.1",
		"mac": "5a:4b:59:38:a8:f1"
	}],
	"real_mem": "10547MB",
	"time_stamp": "2025-01-17T17:36:19+08:00",
	"bw_recv_usage": "0.08MB",
	"bw_sent_usage": "9.26MB",
	"real_mem_usage": 64.37598947107406,
	"ocpf_time_stamp": "2025-01-17T17:37:04+08:00",
	"disk_io_read_usage": "141.82MB",
	"disk_io_write_usage": "5.39MB"
}
	}`
	agent := &model.Agent{}
	err := json.Unmarshal([]byte(agentJson), agent)
	if err != nil {
		log.Fatal(err)
	}
	entity := agent.Health.Metrics(100002, time.Now())
	esClient, err := olivere.GetOrCreateClient(&es.Config{
		Url:      []string{"http://10.16.16.234"},
		Username: "elastic",
		Password: "123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	s := elastic.NewBulkService(esClient).Index("ocpf-metrics-agent-1")
	s.Add(elastic.NewBulkCreateRequest().Doc(entity))
	bulkResponse, err := s.Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if bulkResponse.Errors {
		for _, item := range bulkResponse.Failed() {
			log.Errorf("Error when report metrics %v", item)
		}
	}
}
