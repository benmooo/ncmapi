package necmapi

import (
	"net/http"
)

//说明 : 调用此接口,可获取热门搜索列表
func (api *NeteaseAPI) SearchHotDetail() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["search_hot_detail"], nil, nil)

	return resp, err
}
