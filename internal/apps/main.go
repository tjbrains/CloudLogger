// Copyright 2023 FlexCDN tjbrains@foxmain.com. All rights reserved. Official site: https://tjbrains.com .

package apps

import teaconst "github.com/tjbrains/CloudLogger/internal/const"

func RunMain(f func()) {
	if teaconst.IsMain {
		f()
	}
}
