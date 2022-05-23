package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

// 说明 : 调用此接口 , 传入音乐 id, 可喜欢该音乐
//
// required
// 必选参数 : id: 歌曲 id
//
// optional
// 可选参数 : like: 布尔值 , 默认为 true 即喜欢 , 若传 false, 则取消喜欢
func (api *NeteaseAPI) Like(id int, opts ...hm) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().
		AddCookie("os", "pc").
		AddCookie("appver", "2.7.1.198277").
		SetRealIP("118.88.88.88").
		SetUA(apitypes.UAAndroid)

	data := hm{"alg": "itembased", "time": 3, "like": false}.Merge(opts...).Merge(hm{"trackId": id})

	resp, err := api.Req(http.MethodPost, APIRoutes["like"], data, opt)
	return resp, err
}
