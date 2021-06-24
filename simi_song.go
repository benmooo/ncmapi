package necmapi

import (
	"net/http"
)

// 说明 : 调用此接口 , 传入歌曲 id, 可获得相似歌曲
//
// required
// 必选参数 : id: 歌曲 id
func (api *NeteaseAPI) SimiSong(songId int, opts ...hm) (*APIResponse, error) {
	opt := hm{"limit": 50, "offset": 0}.Merge(opts...).Set("songid", songId)
	resp, err := api.Req(http.MethodPost, APIRoutes["simi_song"], opt, nil)

	return resp, err
}
