package main

import (
	"bytes"
	"github.com/aobco/log"
	"io"
)

func main() {
	marshal := `{"version":1,"time":1742900535078,"tag":"000760L7","cartridges":[{"id":100036,"tape_library_id":100033,"tag":"000761L7","version":0,"model":"Ultrium-7","slot":10,"block_size":0,"capacity":6000000,"write_tar_id":0,"write_block_num":0,"write_total":0,"status":"empty","location":{"status":"on_slot","drive_slot":-1,"drive_device_file":""},"created_time":"2025-03-25T17:43:58+08:00","write_finished":"0001-01-01T00:00:00Z","description":"","creator":""},{"id":100035,"tape_library_id":100033,"tag":"000347L7","version":0,"model":"Ultrium-7","slot":4,"block_size":0,"capacity":6000000,"write_tar_id":0,"write_block_num":0,"write_total":0,"status":"empty","location":{"status":"on_slot","drive_slot":-1,"drive_device_file":""},"created_time":"2025-03-25T17:43:58+08:00","write_finished":"0001-01-01T00:00:00Z","description":"","creator":""},{"id":100034,"tape_library_id":100033,"tag":"000794L7","version":0,"model":"Ultrium-7","slot":3,"block_size":0,"capacity":6000000,"write_tar_id":0,"write_block_num":0,"write_total":0,"status":"empty","location":{"status":"on_slot","drive_slot":-1,"drive_device_file":""},"created_time":"2025-03-25T17:43:58+08:00","write_finished":"0001-01-01T00:00:00Z","description":"","creator":""},{"id":100033,"tape_library_id":100033,"tag":"000760L7","version":0,"model":"Ultrium-7","slot":1,"block_size":0,"capacity":107142,"write_tar_id":0,"write_block_num":0,"write_total":0,"status":"empty","location":{"status":"on_drive","drive_slot":1,"drive_device_file":"/dev/nst0"},"created_time":"2025-03-25T17:43:58+08:00","write_finished":"0001-01-01T00:00:00Z","description":"","creator":""}]}`
	reader := bytes.NewReader([]byte(marshal))
	ba := make([]byte, 0)
	for {
		read, err := reader.Read(ba)
		if err != nil {
			if err == io.EOF {
				log.Infof("write cartridge meta to driver:[%s] done", "test")
				break
			}
			log.Errorf("read cartridge meta error: %+v", err)
			return
		}
		if read == 0 {
			continue
		}
	}
}
