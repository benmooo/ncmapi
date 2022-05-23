package ncmapi

import (
	"net/http"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

// 说明 : 调用此接口,可获取歌手全部歌曲 必选参数 :
// required
// id : 歌手 id
// optional:
// order : hot ,time 按照热门或者时间排序
// limit: 取出歌单数量 , 默认为 50
// offset: 偏移数量 , 用于分页 , 如 :( 评论页数 -1)*50, 其中 50 为 limit 的值
func (api *NeteaseAPI) ArtistSongs(id int, opts ...apitypes.H) (*APIResponse, error) {
	data := defaultArtistSongsPayload().Merge(opts...).Merge(hm{"id": id})
	option := apitypes.DefaultRequestOption().AddCookie("os", "pc")

	resp, err := api.Req(http.MethodPost, APIRoutes["artist_songs"], data, option)
	return resp, err
}

func defaultArtistSongsPayload() hm {
	return hm{
		"private_cloud": true,
		"work_type":     1,
		"order":         "hot",
		"offset":        0,
		"limit":         100,
	}
}
