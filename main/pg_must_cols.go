package main

import (
	"fmt"
	"github.com/aobco/log"
	"time"
	"xorm.io/xorm"
	"xsky.com/ocpf/console/service/license"
	"xsky.com/ocpf/console/service/license/handler"
	"xsky.com/ocpf/pkg/error/xerr"
	"xsky.com/ocpf/pkg/license/constant"
	"xsky.com/ocpf/pkg/pg"
)

func main() {
	pg.Init(&pg.Config{
		Url:     "postgres://console_u:Xsky@123@127.0.0.1:5432/console?sslmode=disable",
		ShowSql: false,
		MaxIdle: 10,
		MaxOpen: 10,
	})
	if err := license.LicenseCache.Init(); err != nil {
		panic(err)
	}
	licenses := license.LicenseCache.ListAllLicenses()
	for _, ls := range licenses {
		if err := pg.Tx(func(session *xorm.Session) error {
			handlers := handler.NewHandlerList()
			expiredTimeStr := ls.ExpiredTime
			oldStatus := ls.Status
			_, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", expiredTimeStr, time.UTC)
			if err != nil {
				// if expired time parse error, set license status to inactive
				log.Warnf("Parse license expired time [%s] error: %+v", expiredTimeStr, err)
				ls.Status = constant.LicenseStatusLapsed
				ls.UpdateReason = fmt.Sprintf("%s, expired time [%s] parse error: %+v", constant.LicenseStatusUpdateReasonLicenseExpiredTimeInvalid, expiredTimeStr, err)
				handlers.Add(handler.HandlerService.GetExpiredUpdateStatusHandler(session, ls))
			}
			handlers.AddAll(handler.HandlerService.GetExpiredHandlers(session, ls, oldStatus))
			if handlers == nil {
				log.Warnf("No handler found for license [%d]", ls.Id)
				return nil
			}
			if err := handlers.HandleAll(); err != nil {
				return xerr.Trace(err)
			}
			return nil
		}); err != nil {
			log.Errorf("Refresh license cache error: %+v", err)
		}
	}
}
