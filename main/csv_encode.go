package main

import (
	"encoding/csv"
	"fmt"
	"github.com/aobco/log"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	csvFile, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("open csv file error: %+v", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	defer func() {
		csvWriter.Flush()
		if err = csvFile.Close(); err != nil {
			log.Errorf("close csv file error: %+v", err)
		}
	}()
	row := make([]string, 0)
	row = append(row, "aaa")
	var longFloat int64 = 1234567890123456789
	longFloatStr := fmt.Sprintf("%v", longFloat)
	row = append(row, fmt.Sprintf("%s 等于 %s", "aaa", longFloatStr))
	if err = csvWriter.Write(row); err != nil {
		log.Fatalf("write csv error: %+v", err)
	}
}
