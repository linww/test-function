package main

import (
	"encoding/json"
	"fmt"
	"github.com/aobco/log"
	"io"
	"os"
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
	//for _, d := range datas {
	//	if d.Time >= 1744214400000 && d.Time < 1744250400000 {
	//		for _, a := range d.Agents {
	//			if a.AgentID != 100144 {
	//				continue
	//			}
	//			total += a.TotalCount
	//			m := fmt.Sprintf("%d %s %s %d", d.Time, d.TimeDisplay, a.AgentName, a.TotalCount)
	//			fmt.Println(m)
	//		}
	//	}
	//}
	for _, d := range datas {
		if d.Time >= 1744128000000 && d.Time < 1744264800000 {
			total += d.TotalCount
			fmt.Println(fmt.Sprintf("%s %s %d", d.ResourceName, d.TimeDisplay, d.TotalCount))
		}
	}
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
