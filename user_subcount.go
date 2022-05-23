package ncmapi

import (
	"net/http"
)

// 说明 : 登录后调用此接口 , 可以获取用户信息
// 获取用户信息 , 歌单，收藏，mv, dj 数量
func (api *NeteaseAPI) UserSubcount() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["user_subcount"], nil, nil)

	return resp, err
}
