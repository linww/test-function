package main

import (
	"github.com/aobco/log"
	"sync"
	"sync/atomic"
	"xsky.com/ocpf/console/dao/alert"
	"xsky.com/ocpf/pkg/alert/model"
	"xsky.com/ocpf/pkg/pg"
)

func main() {
	pg.Init(&pg.Config{
		Url:     "postgres://console_u:Xsky@123@10.16.64.13:5432/console",
		ShowSql: false,
		MaxIdle: 10,
		MaxOpen: 100,
	})
	content := `机台 Agent: 10_16_11_217[100013]

告警 Alert: IP冲突(agent_ip_conflict)

告警详情 Details: 机台[10.16.11.217] Ip冲突(IpConflict)`
	totalAlert := &atomic.Int64{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if totalAlert.Load() > 100000 {
					break
				}
				log.Infof("current alert num: %d", totalAlert.Load())
				if _, err := alert.InfoDao.Insert(pg.Engine, &model.AlertInfo{
					Resource:   "agent",
					Type:       "agent_ip_conflict",
					ResourceId: "100013",
					Content:    content,
					Num:        1,
					Level:      "high",
					Status:     "new",
				}); err != nil {
					log.Errorf("insert alert info error: %v", err)
					continue
				}
				totalAlert.Add(1)
			}
		}()
	}
	wg.Wait()
	log.Infof("insert alert info done")
}
