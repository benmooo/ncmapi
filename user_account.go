package necmapi

import (
	"net/http"
)

// 说明 : 登录后调用此接口 ,可获取用户账号信息
func (api *NeteaseAPI) UserAccount() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["user_account"], nil, nil)

	return resp, err
}
