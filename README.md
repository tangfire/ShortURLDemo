# 项目源地址:
https://github.com/zeromicro/zero-doc/blob/main/doc/shorturl.md

## 完整调用演示

* 在 `shorturl` 目录下执行 `go mod tidy`  整理依赖

* 重新依次启动 redis-server,etcd-server,mysql-server, rpc 服务

```shell
# etcd,redis,mysql 自行根据找教程安装启动
# 启动 rpc 服务
cd rpc/transform
go run transform.go -f etc/transform.yaml
```
* 新开终端启动 api 服务

```shell
cd api
go run shorturl.go -f etc/shorturl-api.yaml
```
* shorten api 调用

  ```shell
  # 新开终端调用
  curl -i "http://localhost:8888/shorten?url=https://go-zero.dev"
  ```

  返回如下：

  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json; charset=utf-8
  Traceparent: 00-fe81053320bb99d1d924021a110765bd-fa915fae41db454d-00
  Date: Thu, 10 Aug 2023 03:32:05 GMT
  Content-Length: 20

  {"shorten":"b0434f"}%
  ```

* expand api 调用

  ```shell
  # shorten 值为上一步返回的值为准，每个人返回的值不一样
  curl -i "http://localhost:8888/expand?shorten=b0434f"
  ```

  返回如下：

  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json; charset=utf-8
  Traceparent: 00-0b11aab486c47a35586d6ed08236afb2-b12387d8cc1e3508-00
  Date: Thu, 10 Aug 2023 03:32:54 GMT
  Content-Length: 29

  {"url":"https://go-zero.dev"}%
  ```
