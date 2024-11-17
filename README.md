# CloudLogger

访问日志接收服务Go语言示例，详情请参考 [访问日志接收服务](https://flexcdn.cn/docs/node/accesslog/endpoint) 。

## 编译

使用 `build/build.sh` 即可编译：
~~~bash
cd ./build
./build.sh linux amd64
~~~

编译的结果在 `dist/` 目录下对应的zip文件，安装时解压并运行 `bin/cloud-logger` 即可。

## 配置

在 `configs/config.yaml` 中可以配置启动的端口号，比如：
~~~yaml
port: 8010
~~~

## 安全

应尽可能地在局域网使用此服务，避免被扫描和攻击。