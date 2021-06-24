package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 传入搜索关键词可获得搜索建议 , 搜索结果同时包含单曲 , 歌手 , 歌单 ,mv 信息
//
// required
// 必选参数 : keywords : 关键词
//
// optional
// 可选参数 : type : 如果传 'mobile' 则返回移动端数据
func (api *NeteaseAPI) SearchSuggest(keywords string, opts ...apitypes.H) (*APIResponse, error) {
	opt := hm{}.Merge(opts...)
	_type := "web"
	if opt["type"] == "mobile" {
		_type = "mobile"
	}
	resp, err := api.Req(http.MethodPost, APIRoutes["search_suggest"]+_type, hm{"s": keywords}, nil)

	return resp, err
}
