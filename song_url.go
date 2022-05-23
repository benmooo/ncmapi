package ncmapi

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
	necmcrypto "github.com/benmooo/ncmapi/crypto"
)

// 说明 : 使用歌单详情接口后 , 能得到的音乐的 id, 但不能得到的音乐 url, 调用此接口, 传入的音乐 id( 可多个 , 用逗号隔开 ),
// 可以获取对应的音乐的 url,未登录状态或者非会员返回试听片段(返回字段包含被截取的正常歌曲的开始时间和结束时间)
//
// required
// 必选参数 : id : 音乐 id
//
// optional
// 可选参数 : br: 码率,默认设置了 999000 即最大码率,如果要 320k 则可设置为 320000,其他类推
func (api *NeteaseAPI) SongUrl(id int, ids ...int) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().
		SetCrypto(necmcrypto.CtryptoEapi).
		AddCookie("os", "pc").
		SetURL("/api/song/enhance/player/url")

	if !api.Client.HasCookie(&BaseURL, "MUSIC_U") {
		token := make([]byte, 16)
		rand.Read(token)
		opt.AddCookie("_ntes_nuid", hex.EncodeToString(token))
	}

	_ids, _ := json.Marshal(append(ids, id))
	data := hm{"ids": string(_ids), "br": 999000}

	resp, err := api.Req(http.MethodPost, APIRoutes["song_url"], data, opt)
	return resp, err
}
