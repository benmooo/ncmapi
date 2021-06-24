package necmapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// 说明 : 调用此接口 , 传入音乐 id(支持多个 id, 用 , 隔开), 可获得歌曲详情
//
// requried
// 必选参数 : ids: 音乐 id, 如 ids=347230
func (api *NeteaseAPI) SongDetail(id int, ids ...int) (*APIResponse, error) {
	var idList []hm
	for _, _id := range append(ids, id) {
		idList = append(idList, hm{"id": _id})
	}
	_idList, _ := json.Marshal(idList)

	resp, err := api.Req(http.MethodPost, APIRoutes["song_detail"], hm{"c": string(_idList)}, nil)
	return resp, err
}

func intSlice2StrSlice(s []int) []string {
	var _s []string
	for _, i := range s {
		_s = append(_s, strconv.Itoa(i))
	}
	return _s
}
