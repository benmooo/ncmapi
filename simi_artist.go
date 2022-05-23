package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

// 说明 : 调用此接口 , 传入歌手 id, 可获得相似歌手
//
// requried
// 必选参数 : id: 歌手 id
func (api *NeteaseAPI) SimiArtist(id int) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption()
	if !api.Client.HasCookie(&BaseURL, "MUSIC_U") {
		opt.AddCookie("MUSIC_A", _anonymous_token)
	}

	resp, err := api.Req(http.MethodPost, APIRoutes["simi_artist"], hm{"artistid": id}, opt)
	return resp, err
}

var _anonymous_token = "8aae43f148f990410b9a2af38324af24e87ab9227c9265627ddd10145db744295fcd8701dc45b1ab8985e142f491516295dd965bae848761274a577a62b0fdc54a50284d1e434dcc04ca6d1a52333c9a"
