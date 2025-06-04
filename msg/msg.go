package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 定义发送的webhook地址
var WebhookURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=66dad66e-6deb-4e57-a232-9e9f06596880"

// 定义发送的body类型和对应字段
type TextMessage struct {
	MsgType string `json:"msgtype"`
	// Text    struct {
	// 	Content string `json:"content"`
	// } `json:"text"`
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
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
