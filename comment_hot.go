package necmapi

import (
	"fmt"
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 传入 type, 资源 id 可获得对应资源热门评论 ( 不需要登录 )
// required
// id : 资源 id
// type: 数字 , 资源类型 , 对应歌曲 , mv, 专辑 , 歌单 , 电台, 视频对应以下类型
// 0: 歌曲
// 1: mv
// 2: 歌单
// 3: 专辑
// 4: 电台
// 5: 视频
//
// optional
// 可选参数 : limit: 取出评论数量 , 默认为 20
// offset: 偏移数量 , 用于分页 , 如 :( 评论页数 -1)*20, 其中 20 为 limit 的值
// before: 分页参数,取上一页最后一项的 time 获取下一页数据(获取超过5000条评论的时候需要用到)
func (api *NeteaseAPI) CommentHot(id int, t apitypes.ResourceTypeCode, opts ...apitypes.H) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := limitOffset(20, 0).Merge(opts...).Set("beforeTime", 0).Set("rid", id)
	_u := replaceAllRouteParams(APIRoutes["comment_hot"], "")
	u := fmt.Sprintf("%s%s%d", _u, t, id)

	resp, err := api.Req(http.MethodPost, u, data, opt)
	return resp, err
}
