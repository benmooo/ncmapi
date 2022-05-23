package ncmapi

import (
	"net/http"
)

// 说明 : 调用此接口 , 可获得已收藏专辑列表
// optional
// limit: 取出数量 , 默认为 25
// offset: 偏移数量 , 用于分页 , 如 :( 页数 -1)*25, 其中 25 为 limit 的值 , 默认 为 0
func (api *NeteaseAPI) AlbumSublist(opts ...hm) (*APIResponse, error) {
	data := limitOffset(25, 0).Merge(hm{"total": true}).Merge(opts...)

	resp, err := api.Req(http.MethodPost, APIRoutes["album_sublist"], data, nil)
	return resp, err
}

// default limit & offset
func limitOffset(limit, offset int) hm {
	return hm{"limit": limit, "offset": offset}
}
