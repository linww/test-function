package main

import (
	"encoding/csv"
	"github.com/aobco/log"
	"os"
)

func main() {
	writeCsv()
}

func writeCsv() {
	csvFile, err := os.Create("/Users/xsky/Downloads/escape.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	data := []string{"1", "2", "3", "4", "5aaa\""}
	if err := csvWriter.Write(data); err != nil {
		log.Fatal(err)
	}
}

func readCsv() {
	csvFile, err := os.Open("/Users/xsky/Downloads/多行查询条件 - 副本 (2).csv")
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(csvFile)
	csvReader.LazyQuotes = true
	for {
		record, err := csvReader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for _, field := range record {
			println(field)
		}
	}
}
