package ncmapi

import (
	"net/http"
)

// 说明 : 登录后调用此接口 , 传入用户 id, 可以获取用户历史评论
//
// requried
// 必选参数 : uid : 用户 id
//
// optional
// 可选参数 :
// limit : 返回数量 , 默认为 10
// time: 上一条数据的time,第一页不需要传,默认为0
func (api *NeteaseAPI) UserCommentHistory(uid int, opts ...hm) (*APIResponse, error) {
	data := hm{
		"compose_reminder":    true,
		"compose_hot_comment": true,
		"limit":               10,
		"time":                0,
	}.Merge(opts...).Merge(hm{"user_id": uid})

	resp, err := api.Req(http.MethodPost, APIRoutes["user_comment_history"], data, nil)
	return resp, err
}
