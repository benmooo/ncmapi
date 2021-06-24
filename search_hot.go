package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口,可获取热门搜索列表(简略)
func (api *NeteaseAPI) SearchHot() (*APIResponse, error) {
	data := hm{"type": 1111}
	opt := apitypes.DefaultRequestOption().SetUA(apitypes.UAIPhone)
	resp, err := api.Req(http.MethodPost, APIRoutes["search_hot"], data, opt)

	return resp, err
}
