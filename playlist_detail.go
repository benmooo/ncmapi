package necmapi

import (
	"net/http"
)

// 说明 : 歌单能看到歌单名字, 但看不到具体歌单内容 , 调用此接口 , 传入歌单 id,
// 可以获取对应歌单内的所有的音乐(未登录状态只能获取不完整的歌单,登录后是完整的)，
// 但是返回的trackIds是完整的，tracks 则是不完整的，
// 可拿全部 trackIds 请求一次 song/detail 接口获取所有歌曲的详情 (https://github.com/Binaryify/NeteaseCloudMusicApi/issues/452)
//
// required
// 必选参数 : id : 歌单 id
//
// optional
// 可选参数 : s : 歌单最近的 s 个收藏者,默认为8
func (api *NeteaseAPI) PlaylistDetail(id int, opts ...hm) (*APIResponse, error) {
	data := hm{"n": 100000, "s": 8}.Merge(opts...).Merge(hm{"id": id})
	resp, err := api.Req(http.MethodPost, APIRoutes["playlist_detail"], data, nil)

	return resp, err
}
