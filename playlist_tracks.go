package necmapi

import (
	"encoding/json"
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 可以添加歌曲到歌单或者从歌单删除某首歌曲 ( 需要登录 )
//
// required
// op: 从歌单增加单曲为 add, 删除为 del
// pid: 歌单 id
// tracks: 歌曲 id,可多个,用逗号隔开
func (api *NeteaseAPI) PlaylistTracks(op string, pid int, tracks []int) (*APIResponse, error) {
	data := hm{"op": op, "pid": pid, "trackIds": intSlice2Json(tracks)}.Merge(hm{"imme": true})
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	resp, err := api.Req(http.MethodPost, APIRoutes["playlist_tracks"], data, opt)

	return resp, err
}

func intSlice2Json(s []int) string {
	b, _ := json.Marshal(s)
	return string(b)
}
