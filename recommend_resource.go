package necmapi

import (
	"net/http"
)

// 说明 : 调用此接口 , 可获得每日推荐歌单 ( 需要登录 )
func (api *NeteaseAPI) RecommendResource() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["recommend_resource"], hm{}, nil)

	return resp, err
}
