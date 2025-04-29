package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"xsky.com/ocpf/pkg/alert/service/dingding"
)

func main() {

	// 2. 上传图片并获取 media_id
	filePath := "./output.png" // 本地图片路径
	token, err := getDingTalkAccessToken("dingqq7dkhss0pk7zivx", "RZg1k_8kFYjNu0g80rv8yB41wZi4lprqR8qdR2WkTIjLHFxupyZsEc2Yq4DnHQoK")
	if err != nil {
		fmt.Println("获取 token 失败:", err)
		return
	}
	mediaId, err := uploadImageToDingTalk(token, filePath)
	if err != nil {
		fmt.Println("上传图片失败:", err)
		return
	}

	// 3. 发送图片消息
	webhookURL := "https://oapi.dingtalk.com/robot/send?access_token=7ea9cdc473ed52efe5550a30f44e8df7fd9aa86fc75c66bb5186f36db4fd7eef"
	targetURL, err := url.Parse(webhookURL)
	if err != nil {
		fmt.Println("解析 URL 失败:", err)
		return
	}

	//	text := `
	//	数据报表推送
	//	![图片](%s)
	//`

	password := "SEC6d62fe2af1e52e55637abeafca64c0dccb45e4cf72c4bb53625c14854d70efdf"
	dingMsg := &dingding.Message{
		MessageType: "markdown",
		WebHookUrl:  webhookURL,
		PassWord:    password,
		Markdown: &dingding.DingTalkNotificationMarkdown{
			Title: "图片",
			Text:  fmt.Sprintf("easydata数据报表推送:\n   ![图片](%s)", mediaId),
		},
	}
	if len(dingMsg.PassWord) > 0 {
		timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
		stringToSign := []byte(timestamp + "\n" + password)

		mac := hmac.New(sha256.New, []byte(password))
		mac.Write(stringToSign) // nolint: errcheck
		signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

		qs := targetURL.Query()
		qs.Set("timestamp", timestamp)
		qs.Set("sign", signature)
		targetURL.RawQuery = qs.Encode()
	}
	body, err := json.Marshal(dingMsg)
	if err != nil {
		fmt.Println("error encoding DingTalk request:", err)
		return
	}

	httpReq, err := http.NewRequest("POST", targetURL.String(), bytes.NewReader(body))
	if err != nil {
		fmt.Println("error creating HTTP request:", err)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("error sending request to DingTalk:", err)
		return
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		fmt.Println("unacceptable response code:", resp.StatusCode)
		return
	}

	var robotResp dingding.DingTalkNotificationResponse
	enc := json.NewDecoder(resp.Body)
	if err = enc.Decode(&robotResp); err != nil {
		fmt.Println("error decoding response from DingTalk:", err)
		return
	}
	if robotResp.IsSuccess() {
	} else {
		fmt.Println("error from DingTalk:", robotResp.ErrorMessage)
		return
	}
}

func getDingTalkAccessToken(appKey, appSecret string) (string, error) {
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", appKey, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.ErrCode != 0 {
		return "", fmt.Errorf("获取 token 失败: %s", result.ErrMsg)
	}
	return result.AccessToken, nil
}

func uploadImageToDingTalk(token, filePath string) (string, error) {
	// 读取本地图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 构造 multipart/form-data 请求
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", err
	}
	writer.Close()

	webHookUrl := fmt.Sprintf("https://oapi.dingtalk.com/media/upload?access_token=%s&type=file", token)
	// 发送上传请求
	targetURL, err := url.Parse(webHookUrl)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", targetURL.String(), body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		MediaId string `json:"media_id"`
		Type    string `json:"type"`
		Created int64  `json:"created_at"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.ErrCode != 0 {
		return "", fmt.Errorf("上传图片失败: %s", result.ErrMsg)
	}
	return result.MediaId, nil
}
