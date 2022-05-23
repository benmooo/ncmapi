package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

// 说明 : 调用此接口 , 可刷新登录状态
func (api *NeteaseAPI) LoginRefresh() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["login_refresh"], apitypes.H{}, nil)

	return resp, err
}
