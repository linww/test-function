package main

import (
	"fmt"
	"regexp"
	"strings"
)

type ScsiDevice struct {
	SCSIID        string // SCSI ID（如 [0:0:0:0]）
	DeviceType    string // 设备类型（如 disk、tape、mediumx）
	Vendor        string // 厂商（如 IBM、QUANTUM）
	Model         string // 型号（如 ULTRIUM-HH7、Scalar i3-i6）
	Revision      string // 固件版本（如 KAH1、225G）
	DeviceFile    string // 设备文件（如 /dev/st0、/dev/st1）
	GenericDevice string // 通用设备文件（如 /dev/sg0）
}

func main() {
	lsscsi := `[0:0:0:0]    cd/dvd  QEMU     QEMU DVD-ROM     2.5+  /dev/sr0   /dev/sg0
[6:0:0:0]    tape    IBM      ULTRIUM-HH7      KAH1  /dev/st0   /dev/sg1
[6:0:0:1]    mediumx QUANTUM  Scalar i3-i6     225G  /dev/sch0  /dev/sg2
[6:0:1:0]    tape    IBM      ULTRIUM-HH7      KAH1  /dev/st1   /dev/sg3`
	// 将输出按行分割
	lines := strings.Split(string(lsscsi), "\n")

	// 定义正则表达式解析每一行
	re := regexp.MustCompile(`^(\[\d+:\d+:\d+:\d+\])\s+(\S+)\s+(\S+)\s+([\S\s]+?)\s+(\S+)\s+(\S+)\s+(\S+)$`)

	// 存储解析结果
	var devices []ScsiDevice

	// 解析每一行
	for _, line := range lines {
		if line == "" {
			continue // 跳过空行
		}

		// 使用正则表达式匹配
		matches := re.FindStringSubmatch(line)
		if len(matches) != 8 {
			fmt.Println("Failed to parse line:", line)
			continue
		}

		// 封装到结构体
		device := ScsiDevice{
			SCSIID:        matches[1],
			DeviceType:    matches[2],
			Vendor:        matches[3],
			Model:         matches[4],
			Revision:      matches[5],
			DeviceFile:    matches[6],
			GenericDevice: matches[7],
		}

		// 添加到结果列表
		devices = append(devices, device)
	}

	// 打印解析结果
	for _, device := range devices {
		fmt.Printf("SCSI ID: %s, \nType: %s, \nVendor: %s, \nModel: %s, \nRevision: %s, \nDevice: %s, \nGeneric: %s\n",
			device.SCSIID, device.DeviceType, device.Vendor, device.Model, device.Revision, device.DeviceFile, device.GenericDevice)
		fmt.Println("================================================")
	}
}
