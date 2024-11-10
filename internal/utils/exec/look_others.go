// Copyright 2023 FlexCDN tjbrains@foxmail.com. All rights reserved. Official site: https://tjbrains.com .
//go:build !linux

package executils

import "os/exec"

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
