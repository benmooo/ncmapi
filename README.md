<h1 align="center">NecmAPI</h1>

[![GoDoc](https://pkg.go.dev/badge/github.com/benmooo/necm-api?status.svg)](https://pkg.go.dev/github.com/benmooo/necm-api?tab=readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmooo/necm-api)](https://goreportcard.com/report/github.com/benmooo/necm-api)

A go package for netease cloud Music API. --Why? Two factors make this repo sense. First, better integration with go projects. If your go prject requires this kindof functions, just bring in this package. Second, 'local app' friendly. 'local' basicly means that an app independent of third api server. Sorry if i express the idea not exactly. Imagine that you have an app consumes necmapi, but you do not have a server to spin [NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) up. A compromise way to handle this is every time your app starts up, spawn a new process to set up the api server locally, which is not a good idea. --When to use? Developing a go projects and won't bother to set up an api server.

---

### Useage

Installation
```sh
$ go get -u github.com/benmooo/necm-api
```

QuickStart

```go
package main

import (
	"fmt"
	necmapi "github.com/benmooo/necm-api"
)

func main() {
	api := necmapi.Default()

	res, _ := api.Search("mota")
	fmt.Println(res)
}
```


### Document

Most of the functions are self documented. If there is some confusion about the params of a funtion requires, figure out [here](https://neteasecloudmusicapi.vercel.app)



### How it works

Necmapi consists of three parts
* api: which is iteself, resolves an APIRequest to retrive data from cache or forwarding an http request.
* client: takes an APIRequst, process it into a http.Request by presenting it with header and encrypt the payload etc. And then send requests to the server and returns the response back.
* store: for caching

### Contribute

If you think this package useful, please do make pull requests.
