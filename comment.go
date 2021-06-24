package necmapi

import (
	"fmt"
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口,可发送评论或者删除评论
//
// New Comment
// required
// t:1 发送, 2 回复
// type: 数字,资源类型,对应歌曲,mv,专辑,歌单,电台,视频对应以下类型
// id:对应资源 id
// content :要发送的内容
// commentId :回复的评论id (回复评论时必填)
//
// Delete Comment
// required
// t:0 删除
// type: 数字,资源类型,对应歌曲,mv,专辑,歌单,电台,视频对应以下类型
// id:对应资源 id
// content :内容 id,可通过 /comment/mv 等接口获取
func (api *NeteaseAPI) Comment(t int, commentType apitypes.ResourceTypeCode, id int, opts ...apitypes.H) (*APIResponse, error) {
	// data := concatPayload(apitypes.H{"t": t, "type": commentType, "id": id}, opts...)
	// option := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	// data := hm{}
	// resp, err := api.Req(http.MethodPost, APIRoutes["comment"], data, nil)
	// return resp, err
	return nil, nil
}

// required
// rid: resource id
// rt:  resource type
// cmt: comment body
func (api *NeteaseAPI) CommentCreate(rid int, rt apitypes.ResourceTypeCode, cmt string) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := hm{"threadId": fmt.Sprintf("%s%d", rt, rid), "content": cmt}
	u := replaceAllRouteParams(APIRoutes["comment"], "add")

	resp, err := api.Req(http.MethodPost, u, data, opt)
	return resp, err
}

// required
// rid: resource id
// rt:  resource type
// reid: the comment id of reply to
// cmt: comment body
func (api *NeteaseAPI) CommentRe(rid int, rt apitypes.ResourceTypeCode, reid int, cmt string) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := hm{"threadId": fmt.Sprintf("%s%d", rt, rid), "content": cmt, "commentId": reid}
	u := replaceAllRouteParams(APIRoutes["comment"], "reply")

	resp, err := api.Req(http.MethodPost, u, data, opt)
	return resp, err
}

// required
// rid: resource id
// rt:  resource type
// cmtid: comment id
func (api *NeteaseAPI) CommentDel(rid int, rt apitypes.ResourceTypeCode, cmtid int) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := hm{"threadId": fmt.Sprintf("%s%d", rt, rid), "commentId": cmtid}
	u := replaceAllRouteParams(APIRoutes["comment"], "delete")

	resp, err := api.Req(http.MethodPost, u, data, opt)
	return resp, err
}
