[中文](README.md) 👈

<h1 align="center">ncmapi</h1>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/benmooo/ncmapi)
[![GoDoc](https://pkg.go.dev/badge/github.com/benmooo/ncmapi?status.svg)](https://pkg.go.dev/github.com/benmooo/ncmapi?tab=readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmooo/ncmapi)](https://goreportcard.com/report/github.com/benmooo/ncmapi)
![GitHub](https://img.shields.io/github/license/benmooo/ncmapi)

A go package for netease cloud Music API. --Why? Two factors make this repo sense. First, better integration with go projects. If your go prject requires this kindof functions, just bring in this package. Second, 'local app' friendly. 'local' basicly means that an app independent of third api server. Sorry if i express the idea not exactly. Imagine that you have an app consumes ncmapi, but you do not have a server to spin [NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) up. A compromise way to handle this is every time your app starts up, spawn a new process to set up the api server locally, which is not a good idea. --When to use? Developing a go projects and won't bother to set up an api server.

---

### Useage

Installation
```sh
$ go get -u github.com/benmooo/ncmapi
```

QuickStart

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

Configuration

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


### Document

Most of the functions are self documented. If there is some confusion about the params of a funtion requires, figure out [here](https://neteasecloudmusicapi.vercel.app)



### How it works

ncmapi consists of three parts
* api: which is iteself, resolves an APIRequest to retrive data from cache or forwarding an http request.
* client: takes an APIRequst, process it into a http.Request by presenting it with header and encrypt the payload etc. And then send requests to the server and returns the response back.
* store: for caching

### Contribute

If you think this package useful, please do make pull requests.
