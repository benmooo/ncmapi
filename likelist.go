package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

// 说明 : 调用此接口 , 传入用户 id, 可获取已喜欢音乐id列表(id数组)
//
// required
// 必选参数 : uid: 用户 id
func (api *NeteaseAPI) LikeList(uid int) (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["likelist"], apitypes.H{"uid": uid}, nil)

	return resp, err
}
