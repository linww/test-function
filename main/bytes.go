package main

import (
	"encoding/binary"
	"math"
)

func main() {
	buf := []byte{0, 0, 0, 0, 255, 255, 255, 255}
	size := int64(buf[0])<<56 | int64(buf[1])<<48 | int64(buf[2])<<40 | int64(buf[3])<<32 |
		int64(buf[4])<<24 | int64(buf[5])<<16 | int64(buf[6])<<8 | int64(buf[7])
	println(size)
	println(math.MaxInt64)
	buf = make([]byte, 8)
	size = 9223372036854775807
	binary.BigEndian.PutUint64(buf, uint64(size))
	println(buf)
}
