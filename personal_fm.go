package ncmapi

import (
	"net/http"
)

// 说明 : 私人 FM( 需要登录 )
func (api *NeteaseAPI) PersonalFm() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["personal_fm"], nil, nil)

	return resp, err
}
