package remotelogs

import (
	teaconst "github.com/tjbrains/CloudLogger/internal/const"
	"github.com/tjbrains/CloudLogger/internal/utils/goman"
	"github.com/tjbrains/TeaGo/Tea"
	"github.com/tjbrains/TeaGo/logs"
	"time"
)

func init() {
	if !teaconst.IsMain {
		return
	}

	// 定期上传日志
	var ticker = time.NewTicker(60 * time.Second)
	if Tea.IsTesting() {
		ticker = time.NewTicker(10 * time.Second)
	}
	goman.New(func() {
		for range ticker.C {
			err := uploadLogs()
			if err != nil {
				logs.Println("[LOG]upload logs failed: " + err.Error())
			}
		}
	})
}

// Debug 打印调试信息
func Debug(tag string, description string) {
	logs.Println("[" + tag + "]" + description)
}

// Println 打印普通信息
func Println(tag string, description string) {
	logs.Println("[" + tag + "]" + description)
}

// Warn 打印警告信息
func Warn(tag string, description string) {
	logs.Println("[" + tag + "]" + description)
}

// WarnServer 打印服务相关警告
func WarnServer(tag string, description string) {
	if Tea.IsTesting() {
		logs.Println("[" + tag + "]" + description)
	}
}

// Error 打印错误信息
func Error(tag string, description string) {
	logs.Println("[" + tag + "]" + description)
}

// ErrorServer 打印服务相关错误信息
func ErrorServer(tag string, description string) {
	if Tea.IsTesting() {
		logs.Println("[" + tag + "]" + description)
	}
}

// ErrorObject 打印错误对象
func ErrorObject(tag string, err error) {
	if err == nil {
		return
	}
}

// 上传日志
func uploadLogs() error {
	// stub
	return nil
}
