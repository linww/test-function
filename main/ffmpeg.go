package main

import (
	"github.com/aobco/log"
	"xsky.com/ocpf/logservice/util"
)

func main() {
	command := `/usr/local/src/ffmpeg-git-amd64/ffmpeg -f lavfi -i testsrc=duration=18000:size=1920x1080:rate=30 -b:v 45M -minrate 45M -maxrate 45M -bufsize 90M -c:v libx264 -preset ultrafast -tune stillimage -pix_fmt yuv420p output_100gb.mp4`
	result, err := util.Sync(command)
	if err != nil {
		panic(err)
	}
	log.Infof(result)
}
