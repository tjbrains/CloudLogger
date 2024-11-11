// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package services

import (
	"context"
	"github.com/tjbrains/CloudLogger/pkg/rpc/pb"
)

type HttpAccessLogService struct {
}

func (this *HttpAccessLogService) CreateHttpAccessLogs(ctx context.Context, req *pb.CreateHttpAccessLogsRequest) (*pb.CreateHttpAccessLogsResponse, error) {
	// req.HttpAccessLogs 为接收到的访问日志列表
	// 这里用你的代码替代
	/**for _, accessLog := range req.HttpAccessLogs {
		fmt.Println(accessLog.RemoteAddr + " [" + accessLog.TimeLocal + "] \"" + accessLog.RequestMethod + " " + accessLog.Scheme + "://" + accessLog.Host + accessLog.RequestURI + " " + accessLog.Proto + "\" " + types.String(accessLog.Status) + " " + types.String(accessLog.BytesSent) + " " + strconv.Quote(accessLog.Referer) + " " + strconv.Quote(accessLog.UserAgent) + " - " + fmt.Sprintf("%.2fms", accessLog.RequestTime*1000) + "\n")
	}**/

	return &pb.CreateHttpAccessLogsResponse{}, nil
}
