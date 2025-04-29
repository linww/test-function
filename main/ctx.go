package main

import (
	"context"
	"github.com/aobco/log"
	"time"
)

func main() {
	i := 0
	for {
		i++
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(5 * time.Second)
			cancel()
		}()
		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Infof("current loop %d, go func 1, cancelled", i)
					return
				default:
					log.Infof("current loop %d, go func 1", i)
					time.Sleep(1 * time.Second)
				}
			}
		}()
		go func() {
			for {
				time.Sleep(1 * time.Second)
				log.Infof("current loop %d, go func 2", i)
			}
		}()
		time.Sleep(2 * time.Second)
	}
}
