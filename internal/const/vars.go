// Copyright 2021 FlexCDN tjbrains@foxmain.com. All rights reserved.

package teaconst

import (
	"os"
	"strings"
)

var (
	IsMain = checkMain()
)

// 检查是否为主程序
func checkMain() bool {
	if len(os.Args) == 1 ||
		(len(os.Args) >= 2 && os.Args[1] == "pprof") {
		return true
	}
	exe, _ := os.Executable()
	return strings.HasSuffix(exe, ".test") ||
		strings.HasSuffix(exe, ".test.exe") ||
		strings.Contains(exe, "___")
}
