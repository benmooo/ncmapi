package ncmapi

import (
	"fmt"
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
	necmcrypto "github.com/benmooo/ncmapi/crypto"
)

// 新版评论接口
// 说明 : 调用此接口 , 传入资源类型和资源id,以及排序方式,可获取对应资源的评论
//
// required
// id : 资源 id, 如歌曲 id,mv id
// type: 数字 , 资源类型 , 对应歌曲 , mv, 专辑 , 歌单 , 电台, 视频对应以下类型
//
// optional
// pageNo:分页参数,第N页,默认为1
// pageSize:分页参数,每页多少条数据,默认20
// sortType: 排序方式,1:按推荐排序,2:按热度排序,3:按时间排序
// cursor: 当sortType为3时且页数不是第一页时需传入,值为上一条数据的time
func (api *NeteaseAPI) CommentNew(id int, t apitypes.ResourceTypeCode, opts ...hm) (*APIResponse, error) {
	option := apitypes.DefaultRequestOption().
		SetCrypto(necmcrypto.CtryptoEapi).
		AddCookie("os", "pc").
		SetURL("/api/v2/resource/comments")
	data := _constructCommentNewPayload(opts...).Set("threadId", fmt.Sprintf("%s%d", t, id))

	resp, err := api.Req(http.MethodPost, APIRoutes["comment_new"], data, option)
	return resp, err
}

func _constructCommentNewPayload(opts ...hm) hm {
	args := hm{}.Merge(opts...)
	pageSize, ok := args["pageSize"]
	if !ok {
		pageSize = 20
	}
	pageNo, ok := args["pageNo"]
	if !ok {
		pageNo = 1
	}
	showInner, ok := args["showInner"]
	if !ok {
		showInner = true
	}
	sortType, ok := args["sortType"]
	if !ok {
		sortType = 1
	}
	cursor, ok := args["cursor"]
	if !ok {
		cursor = 0
	}

	if sortType != 3 {
		cursor = pageSize.(int) * pageNo.(int)
	}
	return hm{
		"pageSize":  pageSize,
		"pageNo":    pageNo,
		"showInner": showInner,
		"sortType":  sortType,
		"cursor":    cursor,
	}
}

// The function above is a reedition of https://github.com/Binaryify/NeteaseCloudMusicApi/module/comment_new.js
// which looks terrible, so here comes the alternative
//
// 说明 : 调用此接口 , 传入资源类型和资源id,以及排序方式,可获取对应资源的评论
//
// required
// id : 资源 id, 如歌曲 id,mv id
// type: 数字 , 资源类型 , 对应歌曲 , mv, 专辑 , 歌单 , 电台, 视频对应以下类型
//
// optional
// pageNo:分页参数,第N页,默认为1
// pageSize:分页参数,每页多少条数据,默认20
// sortType: 排序方式,1:按推荐排序,2:按热度排序,3:按时间排序
// cursor: 当sortType为3时且页数不是第一页时需传入,值为上一条数据的time
func (api *NeteaseAPI) GetComment(id int, t apitypes.ResourceTypeCode, pageSize, pageNo, sortType, cursor int, showInner bool) (*APIResponse, error) {
	option := apitypes.DefaultRequestOption().
		SetCrypto(necmcrypto.CtryptoEapi).
		AddCookie("os", "pc").
		SetURL("/api/v2/resource/comments")

	if sortType != 3 {
		cursor = pageSize * pageNo
	}

	data := hm{
		"pageSize":  pageSize,
		"pageNo":    pageNo,
		"showInner": showInner,
		"sortType":  sortType,
		"cursor":    cursor,
	}.Set("threadId", fmt.Sprintf("%s%d", t, id))

	resp, err := api.Req(http.MethodPost, APIRoutes["comment_new"], data, option)
	return resp, err
}
