package necmapi

import (
	"encoding/json"
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 传入音乐 id, 来源 id，歌曲时间 time，更新听歌排行数据
//
// requried
// 必选参数 :
// id: 歌曲 id
// sourceid: 歌单或专辑 id
//
// optional
// 可选参数 : time: 歌曲播放时间,单位为秒
func (api *NeteaseAPI) Scrobble(id, sourceid int, opts ...apitypes.H) (*APIResponse, error) {
	logs := []hm{
		{
			"action": "play",
			"json": hm{
				"download": 0,
				"end":      "playend",
				"id":       id,
				"sourceId": sourceid,
				"time":     hm{}.Merge(opts...)["time"],
				"type":     "song",
				"wifi":     0,
			},
		},
	}
	_logs, _ := json.Marshal(logs)
	data := hm{"logs": _logs}
	resp, err := api.Req(http.MethodPost, APIRoutes["scrobble"], data, nil)

	return resp, err
}
