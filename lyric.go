package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 传入音乐 id 可获得对应音乐的歌词 ( 不需要登录 )
//
// required
// 必选参数 : id: 音乐 id
func (api *NeteaseAPI) Lyric(id int) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := hm{
		"lv": -1,
		"kv": -1,
		"tv": -1,
	}.Set("id", id)

	resp, err := api.Req(http.MethodPost, APIRoutes["lyric"], data, opt)
	return resp, err
}
