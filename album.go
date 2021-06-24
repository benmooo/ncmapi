package necmapi

import (
	"net/http"
	"strconv"
)

// 说明 : 调用此接口 , 传入专辑 id, 可获得专辑内容
// required
// 必选参数 : id: 专辑 id
func (api *NeteaseAPI) Album(id int) (*APIResponse, error) {
	u := replaceAllRouteParams(APIRoutes["album"], strconv.Itoa(id))

	resp, err := api.Req(http.MethodPost, u, nil, nil)
	return resp, err
}
