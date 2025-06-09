package main

import (
	"fmt"
	"log"
	"net/http"
	"remind/msg"
	"time"

	"github.com/robfig/cron/v3"
)

// 用来判断msg是否发送
var Button bool = true

// 提供前端页面
func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}

// 执行函数的处理程序
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	Button = false

	// 检查HTTP方法
	if r.Method != "POST" {
		// fmt.Println("111")
		http.Error(w, "仅支持POST方法", http.StatusMethodNotAllowed)
		return
	}

	// 调用你的目标函数
	msg.CancelTask()

	// 返回结果
	// w.Write([]byte(result))
}

func main() {

	// 创建cron实例
	c1 := cron.New(cron.WithSeconds())
	// 添加任务(cron表达式:秒 分 时 日 月 周 + cmd)
	_, err1 := c1.AddFunc("0 * 18 * * 1-5", func() {
		if Button == false || (time.Now().Hour() == 18 && time.Now().Minute() >= 30) {
			c1.Stop()
		} else {
			textMsg1 := msg.TextMessage{
				MsgType: "text",
				// Text: struct {
				// 	Content string `json:"content"`
				// }{Content: "this is test"}, // 这是匿名结构体
				Text: msg.Text{
					Content: "卡芙卡来电:记得打卡...",
				},
			}
			textMsg2 := msg.TextMessage{
				MsgType: "image",
				// Text: struct {
				// 	Content string `json:"content"`
				// }{Content: "this is test"}, // 这是匿名结构体
				Image: msg.Image{
					Base64: msg.MsgPhotoBase64(),
					Md5:    msg.MsgPhotoMd5(),
				},
			}
			if err1 := msg.SendMessage(textMsg1); err1 != nil {
				fmt.Printf("发送文本失败!%v\n", err1)
			} else {
				fmt.Println("发送文本成功!")
			}
			time.Sleep(time.Second)
			if err2 := msg.SendMessage(textMsg2); err2 != nil {
				fmt.Printf("发送图片失败!%v\n", err2)
			} else {
				fmt.Println("发送图片成功!")
			}

			// fmt.Printf("定时任务,每两秒执行:%v\n", time.Now().Format("15:04:05"))
		}

	})
	if err1 != nil {
		// fmt.Println("添加任务失败")
		log.Fatal("添加启动任务失败:", err1)
	}

	// 启动定时任务
	c1.Start()

	// 设置定时任务每天开启msg发送
	c2 := cron.New(cron.WithSeconds())
	_, err2 := c2.AddFunc("0 0 9 * * 1-5", func() {
		c1.Start()
		Button = true
	})
	if err2 != nil {
		log.Fatal("添加取消任务失败:", err2)
	}
	c2.Start()

	// 6点再开次任务，防止之前手欠提前误点
	c3 := cron.New(cron.WithSeconds())
	_, err3 := c3.AddFunc("0 0 18 * * 1-5", func() {
		Button = true
		c1.Start()
	})
	if err3 != nil {
		log.Fatal("添加取消任务失败:", err2)
	}
	c3.Start()

	// 配置静态路径，用于html的图片路径识别
	fs := http.FileServer(http.Dir("./web")) // 假设图片在项目根目录的 public 文件夹
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// 设置路由
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/shutdown", shutdownHandler)
	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /", serveIndex)
	// mux.HandleFunc("POST /shutdown", shutdownHandler)

	// 启动HTTP服务器
	fmt.Println("服务器启动，监听 :8080")
	fmt.Println("访问 http://127.0.0.1:8080 点击按钮执行函数")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("服务器启动失败: ", err)
	}

	// 防止程序退出
	select {}
}
