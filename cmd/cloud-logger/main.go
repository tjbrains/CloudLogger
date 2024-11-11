// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package main

import (
	"github.com/tjbrains/CloudLogger/internal/apps"
	teaconst "github.com/tjbrains/CloudLogger/internal/const"
	"github.com/tjbrains/CloudLogger/internal/nodes"
)

func main() {
	var app = apps.NewAppCmd().
		Version(teaconst.Version).
		Product(teaconst.ProductName + "-Logger").
		Usage(teaconst.ProcessName + " [-v|start|stop|restart|status]")

	app.Run(func() {
		var node = nodes.NewNode()
		node.Start()
	})
}
