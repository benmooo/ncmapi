package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 登录后调用此接口 , 传入用户 id, 可获取用户播放记录
//
// requred
// 必选参数 : uid : 用户 id
//
// optional
// 可选参数 : type : type=1 时只返回 weekData, type=0 时返回 allData
func (api *NeteaseAPI) UserRecord(uid int, opts ...apitypes.H) (*APIResponse, error) {
	data := hm{"type": 1}.Merge(opts...).Merge(hm{"uid": uid})

	resp, err := api.Req(http.MethodPost, APIRoutes["user_record"], data, nil)
	return resp, err
}
