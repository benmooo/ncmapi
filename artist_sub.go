package necmapi

import (
	"fmt"
	"net/http"
)

// 说明 : 调用此接口,可收藏歌手
// required
// id : 歌手 id
// t:操作,1 为收藏,其他为取消收藏
func (api *NeteaseAPI) ArtistSub(id, t int) (*APIResponse, error) {
	op := "sub"
	if t != 1 {
		op = "unsub"
	}
	data := hm{"artistId": id, "artistIds": fmt.Sprintf("[%d]", id)}
	u := replaceAllRouteParams(APIRoutes["artist_sub"], op)

	resp, err := api.Req(http.MethodPost, u, data, nil)
	return resp, err
}
