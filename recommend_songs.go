package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 可获得每日推荐歌曲 ( 需要登录 )
func (api *NeteaseAPI) RecommendSongs() (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "ios")
	resp, err := api.Req(http.MethodPost, APIRoutes["recommend_songs"], nil, opt)

	return resp, err
}
