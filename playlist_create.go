package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 传入歌单名字可新建歌单
//
// required
// 必选参数 : name : 歌单名
//
// optional
// privacy : 是否设置为隐私歌单，默认否，传'10'则设置成隐私歌单
// type : 歌单类型,默认'NORMAL',传 'VIDEO'则为视频歌单
func (api *NeteaseAPI) PlaylistCreate(name string, opts ...apitypes.H) (*APIResponse, error) {
	option := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := hm{"type": "NORMAL", "privacy": 0}.Merge(opts...).Set("name", name)

	resp, err := api.Req(http.MethodPost, APIRoutes["playlist_create"], data, option)
	return resp, err
}
