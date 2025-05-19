package main

import (
	"github.com/aobco/log"
	"github.com/vjeantet/grok"
)

func main() {
	result, err := ParseGrok("", "cifar-10-batches-py/ApPFAOIAfterTest/A401/S0931024000D00/20250516/D4VHET1047F000126W+5K02K03XP10/S00001_C03_P02_L0_PI139_G1_M1_Y1_20250516114831401_D4VHET1047F000126W+5K02K03XP10_NY.jpg", []string{"(?:.*[\\/](?:[^_]*_){9}(?P<sn>[^_]+).*.jpg)", "(.*[\\/](?P<sn>.*)\\.jpg)"})
	if err != nil {
		log.Errorf("ParseGrok error: %v", err)
		return
	}
	log.Infof("result: %v", result)
}

func ParseGrok(bucket, key string, patterns []string) (map[string]string, error) {
	g, err := grok.New()
	if err != nil {
		log.Fatalf("new grok error: %v", err)
	}
	result := make(map[string]string)
	for _, p := range patterns {
		parseMap, err := g.Parse(p, key)
		if err != nil {
			log.Errorf("bucket: [%s], key: [%s], ParseGrok error: %v", bucket, key, err)
			return nil, err
		}
		for k, v := range parseMap {
			result[k] = v
		}
	}
	return result, nil
}
