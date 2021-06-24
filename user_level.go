package necmapi

import (
	"net/http"
)

// 说明 : 登录后调用此接口 , 可以获取用户等级信息,包含当前登录天数,听歌次数,下一等级需要的登录天数和听歌次数,当前等级进度
// 对应 https://music.163.com/#/user/level
func (api *NeteaseAPI) UserLevel() (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["user_level"], nil, nil)

	return resp, err
}
