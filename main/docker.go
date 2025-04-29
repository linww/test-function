package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//var commandResult string
	loginCmd := exec.Command("docker", "login", "10.16.64.86:7508", "-u", "admin", "-p", "Xsky@123")
	stdout, err := loginCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error logging in: %v: %s", err, stdout)
		return
	}
	//// Load Docker image
	//loadCmd := exec.Command("docker", "load", "-i", "/Users/xsky/Desktop/working/3rd.plugin/build/unzip/unzip-20241212144445.tar")
	//stdout, err = loadCmd.CombinedOutput()
	//if err != nil {
	//	fmt.Printf("Error load image: %v: %s", err, stdout)
	//	return
	//}
	//commandResult = string(stdout)
	//var imageName string
	//var tag string
	//if strings.Contains(commandResult, "Loaded image:") {
	//	var imageNameTag string
	//	commandResult = strings.Trim(commandResult, "\n")
	//	parts := strings.Split(commandResult, ": ")
	//	if len(parts) > 1 {
	//		imageNameTag = parts[1]
	//	}
	//	imageNameTagParts := strings.Split(imageNameTag, ":")
	//	if len(imageNameTagParts) > 1 {
	//		imageName = imageNameTagParts[0]
	//		tag = imageNameTagParts[1]
	//	}
	//}
	//var pushImageName string
	//parts := strings.Split(imageName, "/")
	//if len(parts) > 1 {
	//	parts[0] = "ocpf"
	//	pushImageName = strings.Join(parts, "/")
	//}
	//// Tag Docker image
	//tagCmd := exec.Command("docker", "tag", fmt.Sprintf("%s:%s", imageName, tag), fmt.Sprintf("%s/%s:%s", "10.16.53.16:7508", pushImageName, tag))
	//stdout, err = tagCmd.CombinedOutput()
	//if err != nil {
	//	fmt.Printf("Error tagging image: %v: %s", err, stdout)
	//	return
	//}
	//// Push Docker image
	//pushCmd := exec.Command("docker", "push", fmt.Sprintf("%s/%s:%s", "10.16.53.16:7508", pushImageName, tag))
	//stdout, err = pushCmd.CombinedOutput()
	//if err != nil {
	//	fmt.Printf("Error pushing image: %v: %s", err, stdout)
	//	return
	//}
}
