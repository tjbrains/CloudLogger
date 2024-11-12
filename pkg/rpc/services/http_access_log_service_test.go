// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package services_test

import (
	"context"
	"fmt"
	"github.com/tjbrains/CloudLogger/pkg/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"runtime"
	"strconv"
	"testing"
	"time"
)

const rpcAddr = "192.168.2.61:8010"

func TestHttpAccessLogService_CreateHttpAccessLogs(t *testing.T) {
	client, err := grpc.NewClient(rpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	var service = pb.NewHttpAccessLogServiceClient(client)
	_, err = service.CreateHttpAccessLogs(context.Background(), &pb.CreateHttpAccessLogsRequest{
		HttpAccessLogs: []*pb.HttpAccessLog{
			{
				RequestId:     "1",
				NodeId:        1,
				RemoteAddr:    "127.0.0.1",
				Host:          "example.com",
				RequestMethod: "GET",
				RequestURI:    "/html",
				Status:        200,
				Proto:         "HTTP/1.1",
				Scheme:        "http",
				BytesSent:     1024,
				UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
				TimeLocal:     "10/Nov/2024:15:24:06 +0800",
			},
			{
				RequestId:     "2",
				NodeId:        2,
				RemoteAddr:    "127.0.0.1",
				Host:          "example.com",
				RequestMethod: "GET",
				RequestURI:    "/html",
				Status:        200,
				Proto:         "HTTP/1.1",
				Scheme:        "http",
				BytesSent:     1024,
				UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
				TimeLocal:     "10/Nov/2024:15:24:06 +0800",
			},
			{
				RequestId:     "3",
				NodeId:        3,
				RemoteAddr:    "127.0.0.1",
				Host:          "example.com",
				RequestMethod: "GET",
				RequestURI:    "/html",
				Status:        200,
				Proto:         "HTTP/1.1",
				Scheme:        "http",
				BytesSent:     1024,
				UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
				TimeLocal:     "10/Nov/2024:15:24:06 +0800",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestHttpAccessLogService_CreateHttpAccessLogs_Batch(t *testing.T) {
	client, err := grpc.NewClient(rpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	client.Connect()

	var before = time.Now()

	const count = 5_000
	const loop = 10

	for j := 0; j < loop; j ++ {
		var service = pb.NewHttpAccessLogServiceClient(client)
		var accessLogs = make([]*pb.HttpAccessLog, 0, count)
		for i := 0; i < count; i++ {
			accessLogs = append(accessLogs, &pb.HttpAccessLog{
				RequestId:     strconv.Itoa(i),
				NodeId:        1,
				RemoteAddr:    "127.0.0.1",
				Host:          "example.com",
				RequestMethod: "GET",
				RequestURI:    "/html",
				Status:        200,
				Proto:         "HTTP/1.1",
				Scheme:        "http",
				BytesSent:     1024,
				UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
				TimeLocal:     "10/Nov/2024:15:24:06 +0800",
			})
		}
		_, createErr := service.CreateHttpAccessLogs(context.Background(), &pb.CreateHttpAccessLogsRequest{
			HttpAccessLogs: accessLogs,
		})
		if createErr != nil {
			t.Fatal(createErr)
		}
	}

	var costSeconds = time.Since(before).Seconds()

	t.Log("cost:", fmt.Sprintf("%.3fs", costSeconds), "qps:", fmt.Sprintf("%.2f/s", count * loop/costSeconds))
}

func BenchmarkHttpAccessLogService_CreateHttpAccessLogs(b *testing.B) {
	runtime.GOMAXPROCS(32)

	client, err := grpc.NewClient(rpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Fatal(err)
	}

	client.Connect()

	b.ResetTimer()

	var service = pb.NewHttpAccessLogServiceClient(client)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			const count = 3
			var accessLogs = make([]*pb.HttpAccessLog, 0, count)
			for i := 0; i < count; i++ {
				accessLogs = append(accessLogs, &pb.HttpAccessLog{
					RequestId:     strconv.Itoa(i),
					NodeId:        1,
					RemoteAddr:    "127.0.0.1",
					Host:          "example.com",
					RequestMethod: "GET",
					RequestURI:    "/html",
					Status:        200,
					Proto:         "HTTP/1.1",
					Scheme:        "http",
					BytesSent:     1024,
					UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
					TimeLocal:     "10/Nov/2024:15:24:06 +0800",
				})
			}
			_, createErr := service.CreateHttpAccessLogs(context.Background(), &pb.CreateHttpAccessLogsRequest{
				HttpAccessLogs: accessLogs,
			})
			if createErr != nil {
				b.Fatal(createErr)
			}
		}
	})
}

func BenchmarkHttpAccessLogService_CreateHttpAccessLogs_Batch(b *testing.B) {
	runtime.GOMAXPROCS(32)

	client, err := grpc.NewClient(rpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Fatal(err)
	}

	client.Connect()

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			const count = 2000
			var service = pb.NewHttpAccessLogServiceClient(client)
			var accessLogs = make([]*pb.HttpAccessLog, 0, count)
			for i := 0; i < count; i++ {
				accessLogs = append(accessLogs, &pb.HttpAccessLog{
					RequestId:     strconv.Itoa(i),
					NodeId:        1,
					RemoteAddr:    "127.0.0.1",
					Host:          "example.com",
					RequestMethod: "GET",
					RequestURI:    "/html",
					Status:        200,
					Proto:         "HTTP/1.1",
					Scheme:        "http",
					BytesSent:     1024,
					UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
					TimeLocal:     "10/Nov/2024:15:24:06 +0800",
				})
			}
			_, createErr := service.CreateHttpAccessLogs(context.Background(), &pb.CreateHttpAccessLogsRequest{
				HttpAccessLogs: accessLogs,
			})
			if createErr != nil {
				b.Fatal(createErr)
			}
		}
	})
}
