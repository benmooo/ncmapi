package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
	necmcrypto "github.com/benmooo/ncmapi/crypto"
)

// 说明 : 调用此接口 , 可获取默认搜索关键词
func (api *NeteaseAPI) SearchDefault() (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().SetCrypto(necmcrypto.CtryptoEapi).SetURL("/api/search/defaultkeyword/get")

	resp, err := api.Req(http.MethodPost, APIRoutes["search_default"], nil, opt)

	return resp, err
}
