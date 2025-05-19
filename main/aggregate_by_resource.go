package main

import (
	"encoding/json"
	"fmt"
	"github.com/aobco/log"
	"io"
	"os"
	"time"
	"xsky.com/ocpf/pkg/constant"
)

func main() {
	file, err := os.Open("/Users/xsky/Desktop/aggragate-by-resource.txt")
	if err != nil {
		log.Panic("open file error, %v", err)
	}
	dataBytes, err := io.ReadAll(file)
	if err != nil {
		log.Panic("read all error, %v", err)
	}
	datas := make([]*data, 0)
	err = json.Unmarshal(dataBytes, &datas)
	if err != nil {
		panic(err)
	}
	total := int64(0)
	for _, d := range datas {
		if d.Time >= 1746360000000 && d.Time < 1746403200000 {
			for _, a := range d.Agents {
				if a.AgentID != 100067 {
					continue
				}
				total += a.TotalCount
				localTime, _ := time.ParseInLocation(constant.EsDefaultGoDateFormat, d.TimeDisplay, time.UTC)
				m := fmt.Sprintf("%d %s %s %d", d.Time, localTime.Local(), a.AgentName, a.TotalCount)
				fmt.Println(m)
			}
		}
	}
	//for _, d := range datas {
	//	if d.Time >= 1746360000000 && d.Time < 1746403200000 {
	//		total += d.TotalCount
	//		fmt.Println(fmt.Sprintf("%s %s %d", d.ResourceName, d.TimeDisplay, d.TotalCount))
	//	}
	//}
	fmt.Println("total:", total)
}

type data struct {
	Time         int64    `json:"time"`
	TimeDisplay  string   `json:"time_display"`
	Agents       []*agent `json:"agents"`
	TotalCount   int64    `json:"total_count"`
	ResourceName string   `json:"resource_name"`
}

type agent struct {
	AgentID    int    `json:"agent_id"`
	AgentName  string `json:"agent_name"`
	TotalCount int64  `json:"total_count"`
	TotalSize  int    `json:"total_size"`
}
