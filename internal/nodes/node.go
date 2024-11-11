// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package nodes

import (
	"bytes"
	"github.com/iwind/gosock/pkg/gosock"
	"github.com/tjbrains/CloudLogger/internal/remotelogs"
	"github.com/tjbrains/CloudLogger/internal/utils"
	"github.com/tjbrains/CloudLogger/internal/utils/goman"
	"github.com/tjbrains/TeaGo/Tea"
	"github.com/tjbrains/TeaGo/logs"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

type Node struct {
	sock *gosock.Sock
}

func NewNode() *Node {
	return &Node{
		sock: utils.NewSock(),
	}
}

func (this *Node) Start() {
	remotelogs.Println("NODE", "starting ...")

	_ = os.Setenv("GODEBUG", "netdns=go")

	// 处理异常
	this.handlePanic()

	// 监听signal
	this.listenSignals()

	// 启动rpc server
	err := NewServer().Listen()
	if err != nil {
		remotelogs.Error("NODE", "listening server failed: "+err.Error())
		return
	}

	// hold住进程
	select {}
}

// 监听一些信号
func (this *Node) listenSignals() {
	var queue = make(chan os.Signal, 8)
	signal.Notify(queue, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)
	goman.New(func() {
		for range queue {
			time.Sleep(100 * time.Millisecond)
			os.Exit(0)
			return
		}
	})
}

// 处理异常
func (this *Node) handlePanic() {
	// 如果是在前台运行就直接返回
	backgroundEnv, _ := os.LookupEnv("CloudBackground")
	if backgroundEnv != "on" {
		return
	}

	var panicFile = Tea.Root + "/logs/panic.log"

	// 分析panic
	data, err := os.ReadFile(panicFile)
	if err == nil {
		var index = bytes.Index(data, []byte("panic:"))
		if index >= 0 {
			remotelogs.Error("NODE", "系统错误，请上报给开发者: "+string(data[index:]))
		}
	}

	fp, err := os.OpenFile(panicFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		logs.Println("NODE", "open 'panic.log' failed: "+err.Error())
		return
	}
	err = debug.SetCrashOutput(fp, debug.CrashOptions{})
	if err != nil {
		logs.Println("NODE", "write to 'panic.log' failed: "+err.Error())
	}
}
