package necmapi

import (
	"net/http"
)

// 说明 : 登录后调用此接口 , 传入用户 id, 可以获取用户歌单
//
// required
// 必选参数 : uid : 用户 id
//
// optional
// 可选参数 :
// limit : 返回数量 , 默认为 30
// offset : 偏移数量，用于分页 , 如 :( 页数 -1)*30, 其中 30 为 limit 的值 , 默认为 0
func (api *NeteaseAPI) UserPlaylist(uid int, opts ...hm) (*APIResponse, error) {
	data := hm{
		"limit":        30,
		"offset":       0,
		"includeVideo": true,
	}.Merge(opts...).Merge(hm{"uid": uid})

	resp, err := api.Req(http.MethodPost, APIRoutes["user_playlist"], data, nil)
	return resp, err
}
