package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

// 说明 : 调用此接口 , 传入签到类型 ( 可不传 , 默认安卓端签到 ), 可签到 ( 需要登录 ), 其中安卓端签到可获得 3 点经验 , web/PC 端签到可获得 2 点经验
//
// optional
// 可选参数 : type: 签到类型 , 默认 0, 其中 0 为安卓端签到 ,1 为 web/PC 签到
func (api *NeteaseAPI) DailySignin(opts ...apitypes.H) (*APIResponse, error) {
	data := hm{"type": 0}.Merge(opts...)
	resp, err := api.Req(http.MethodPost, APIRoutes["daily_signin"], data, nil)

	return resp, err
}
