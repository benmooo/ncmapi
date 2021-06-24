package necmapi

import (
	"net/http"
)

// 说明 : 调用此接口 , 可退出登录
func (api *NeteaseAPI) Logout() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["logout"], hm{}, nil)

	return resp, err
}
