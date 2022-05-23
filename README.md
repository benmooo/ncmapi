[EN](README.en.md) 👈

<h1 align="center">ncmapi</h1>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/benmooo/ncmapi)
[![GoDoc](https://pkg.go.dev/badge/github.com/benmooo/ncmapi?status.svg)](https://pkg.go.dev/github.com/benmooo/ncmapi?tab=readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmooo/ncmapi)](https://goreportcard.com/report/github.com/benmooo/ncmapi)
![GitHub](https://img.shields.io/github/license/benmooo/ncmapi)


Go版本网易云音乐API. 为什么会有这个项目? 两点, 一来是没有合适的服务器来运行[NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi); 二来就是可以更好的整合进其他go项目.

---

### 用法

安装
```sh
$ go get -u github.com/benmooo/ncmapi
```

使用
```go
package main

import (
	"fmt"
	ncmapi "github.com/benmooo/ncmapi"
)

func main() {
	api := ncmapi.Default()

	res, _ := api.Search("mota")
	fmt.Println(res)
}
```

配置
```go
func main() {
	api := ncmapi.New(
		&ncmapi.NeteaseAPIConfig{
			CacheDefaultExpiration: time.Minute * 1,
			CacheCleanupInterval:   time.Minute * 2,
			PreserveCookies:        true,
		},
	)

	// ...
}
```


### 文档

大多数函数都有很好注释, 如果发现有地方不太清楚, 可以参考 [这里](https://neteasecloudmusicapi.vercel.app)



### 原理

ncmapi 主要由三部分组成
* api: 接受APIRequest, 根据这个请求的id决定发送http请求或者从缓存中取数据.
* client: 模拟客户端, 负责点缀APIRequest, 加密payload, 然后发送http请求, 解密response等等.
* store: 用于缓存

### 贡献

欢迎PR.
