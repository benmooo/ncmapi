<h1 align="center">NecmAPI</h1>

A go package for netease cloud Music API. --Why? Two factors make this repo sense. First, better integration with go projects. If your go prject requires this kindof functions, just bring in this package. Second, 'local app' friendly. 'local' basicly means that an app independent of third api server. Sorry if i express the idea not exactly. Imagine that you have an app consumes necmapi, but you do not have a server to spin [NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi) up. A compromise way to handle this is every time your app starts up, spawn a new process to set up the api server locally, which is not a good idea. --When to use? Developing a go projects and won't bother to set a api server.

---

### Useage
```
api := necmapi.Default()
res := api.Search("mota")

fmt.Println(res.Data)
```

### Document
// todo


### How it works
// working flow


