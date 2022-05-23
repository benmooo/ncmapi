package ncmapi

import (
	"net/http"
)

// 说明 : 调用此接口,可收藏/取消收藏专辑
// required
// id : 专辑 id
// t : 1 为收藏,其他为取消收藏
func (api *NeteaseAPI) AlbumSub(id, t int) (*APIResponse, error) {
	op := "unsub"
	if t == 1 {
		op = "sub"
	}
	data := hm{"id": id}
	u := replaceAllRouteParams(APIRoutes["album_sub"], op)

	resp, err := api.Req(http.MethodPost, u, data, nil)
	return resp, err
}
