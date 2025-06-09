package msg

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 定义发送的webhook地址
var WebhookURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=8c807f42-3cb4-450b-87d9-5c64d824ce06"

// 定义发送的body类型和对应字段
type TextMessage struct {
	MsgType string `json:"msgtype"`
	// Text    struct {
	// 	Content string `json:"content"`
	// } `json:"text"`
	Text  Text  `json:"text"`
	Image Image `json:"image"`
}

type Image struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

type Text struct {
	Content string `json:"content"`
}

// 定义发送的msg中的图片Md5
func MsgPhotoMd5() (Md5 string) {
	filePath := "./web/photo/kafka.jpeg"
	fileDate, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败:%v\n", err)
		return err.Error()
	}

	// 计算md5哈希
	md5Hash := md5.Sum(fileDate)
	md5String := fmt.Sprintf("%x", md5Hash)

	return md5String
}

// 定义发送的msg中的图片Base64
func MsgPhotoBase64() (Base64 string) {
	filePath := "./web/photo/kafka.jpeg"
	fileDate, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件失败:%v\n", err)
		return err.Error()
	}

	// 转换base64
	base64String := base64.StdEncoding.EncodeToString(fileDate)

	return base64String
}

// 定义message发送的方法，通过webhook地址发送
func SendMessage(message interface{}) error {
	// struct序列化成json
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("JSON编码失败:%w", err)
	}
	// 发送消息
	resp, err := http.Post(WebhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("HTTP请求失败..%w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("服务返回错误状态码:%v", resp.StatusCode)
	}

	return nil
}

// 取消今天的任务

func CancelTask() {
	fmt.Println("任务取消成功")
}
