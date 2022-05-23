package ncmapi

import (
	"net/http"
	"regexp"
	"strconv"
)

// 说明 : 登录后调用此接口 , 传入用户 id, 可以获取用户电台
//
// required
// 必选参数 : uid : 用户 id
func (api *NeteaseAPI) UserDj(uid int, opts ...hm) (*APIResponse, error) {
	u := replaceAllRouteParams(APIRoutes["user_dj"], strconv.Itoa(uid))

	data := hm{"limit": 30, "offset": 0}.Merge(opts...)
	resp, err := api.Req(http.MethodPost, u, data, nil)

	return resp, err
}

func replaceAllRouteParams(u, repl string) string {
	reg := regexp.MustCompile(`\${.*}`)
	return reg.ReplaceAllString(u, repl)
}
