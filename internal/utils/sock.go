// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package utils

import (
	"github.com/iwind/gosock/pkg/gosock"
	teaconst "github.com/tjbrains/CloudLogger/internal/const"
)

var sock *gosock.Sock

// NewSock get current process sock
func NewSock() *gosock.Sock {
	if sock != nil {
		return sock
	}

	return gosock.NewTmpSock(teaconst.ProcessName)
}
