package ncmapi

import (
	"fmt"
	"math/rand"
	"net/http"
)

// 说明 : 调用此接口 , 传入音乐 id, 可把该音乐从私人 FM 中移除至垃圾桶
//
// required
// id: 歌曲 id
func (api *NeteaseAPI) FmTrash(id int, opts ...hm) (*APIResponse, error) {
	data := hm{"songId": id}
	time := hm{"time": rand.Intn(20) + 10}.Merge(opts...)["time"]
	u := fmt.Sprintf(APIRoutes["fm_trash"], id, time)

	resp, err := api.Req(http.MethodPost, u, data, nil)
	return resp, err
}
