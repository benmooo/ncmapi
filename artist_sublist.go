package ncmapi

import (
	"net/http"
)

//说明 : 调用此接口,可获取收藏的歌手列表
func (api *NeteaseAPI) ArtistSublist(opts ...hm) (*APIResponse, error) {
	data := limitOffset(25, 0).Merge(opts...).Merge(hm{"total": true})

	resp, err := api.Req(http.MethodPost, APIRoutes["artist_sublist"], data, nil)
	return resp, err
}
