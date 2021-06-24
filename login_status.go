package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口,可获取登录状态
func (api *NeteaseAPI) LoginStatus() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["login_status"], apitypes.H{}, nil)

	return resp, err
}
