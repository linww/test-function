package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aobco/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	corpID     = "wwa73697d0bce60343"
	corpSecret = "S67sdEg5zC7_5Kbb9Ea5courRJRzlameHX4H2IQRjNY"
	agentID    = 1000027 // 你的应用AgentID
)

func main() {
	// 1. 获取AccessToken
	token, err := getWeComToken(corpID, corpSecret)
	if err != nil {
		panic(err)
	}

	// 2. 上传图片获取media_id
	mediaID, err := uploadImage(token, "./output.png")
	if err != nil {
		panic(err)
	}

	// 3. 发送图片消息
	err = sendImageMessage(token, mediaID, "262", "LiBo") // 接收人UserID列表
	if err != nil {
		panic(err)
	}
	fmt.Println("图片消息发送成功！")
}

// 获取AccessToken
func getWeComToken(corpID, corpSecret string) (string, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpID, corpSecret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

// 上传图片
func uploadImage(token, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	writer.Close()

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=image", token)
	resp, err := http.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Infof("resultBytes: %s", string(resultBytes))

	var result struct {
		MediaID string `json:"media_id"`
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	if err := json.NewDecoder(bytes.NewReader(resultBytes)).Decode(&result); err != nil {
		return "", err
	}
	return result.MediaID, nil
}

// 发送图片消息
func sendImageMessage(token, mediaID, toParty, toUser string) error {
	msg := map[string]interface{}{
		"toparty": toParty,
		"touser":  toUser,
		"msgtype": "image",
		"agentid": agentID,
		"image": map[string]string{
			"media_id": mediaID,
		},
	}

	//msg := map[string]interface{}{
	//	"touser":  toUser,
	//	"msgtype": "markdown",
	//	"agentid": agentID,
	//	"markdown": map[string]string{
	//		"content": fmt.Sprintf("easydata数据报表推送:\n   ![图片](%s)", mediaID), // 通过media_id引用图片
	//	},
	//}

	body, _ := json.Marshal(msg)
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Infof("resultBytes: %s", string(resultBytes))
	defer resp.Body.Close()

	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	return json.NewDecoder(bytes.NewReader(resultBytes)).Decode(&result)
}
