package necmapi

import (
	"net/http"
)

// 说明 : 调用此接口,可获取歌手热门50首歌曲
// required
// id : 歌手 id
func (api *NeteaseAPI) ArtistTopSong(id int) (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["artist_top_song"], hm{"id": id}, nil)

	return resp, err
}
