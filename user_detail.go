package necmapi

import (
	"net/http"
	"strconv"
)

// 说明 : 登录后调用此接口 , 传入用户 id, 可以获取用户详情
//
// required
// 必选参数 : uid : 用户 id
func (api *NeteaseAPI) UserDetail(uid int) (*APIResponse, error) {
	u := replaceAllRouteParams(APIRoutes["user_detail"], strconv.Itoa(uid))

	resp, err := api.Req(http.MethodPost, u, nil, nil)
	return resp, err
}
