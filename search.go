package necmapi

import (
	"net/http"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 调用此接口 , 传入搜索关键词可以搜索该音乐 / 专辑 / 歌手 / 歌单 / 用户 , 关键词可以多个 , 以空格隔开 ,
// 如 " 周杰伦 搁浅 "( 不需要登录 ), 搜索获取的 mp3url 不能直接用 , 可通过 /song/url 接口传入歌曲 id 获取具体的播放链接
//
// required
// 必选参数 : keywords : 关键词
//
// optional
// 可选参数 : limit : 返回数量 , 默认为 30 offset : 偏移数量，用于分页 , 如 : 如 :( 页数 -1)*30, 其中 30 为 limit 的值 , 默认为 0
// type: 搜索类型；默认为 1 即单曲 , 取值意义 : 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频, 1018:综合
func (api *NeteaseAPI) Search(keyword string, opts ...hm) (*APIResponse, error) {
	return api.search(keyword, "search", opts...)
}

func (api *NeteaseAPI) CloudSearch(keyword string, opts ...hm) (*APIResponse, error) {
	return api.search(keyword, "cloudsearch", opts...)
}

func (api *NeteaseAPI) search(keyword string, routename string, opts ...hm) (*APIResponse, error) {
	payload := defaultSearchPayload().Merge(opts...).Set("s", keyword)
	return api.Req(http.MethodPost, APIRoutes[routename], payload, nil)

}

func defaultSearchPayload() hm {
	return hm{
		"s":      "*",
		"type":   apitypes.SearchTypeSong,
		"limit":  30,
		"offset": 0,
	}
}

func concatPayload(host hm, payloads ...hm) hm {
	for _, payload := range payloads {
		for k, v := range payload {
			host[k] = v
		}
	}
	return host
}

func (api *NeteaseAPI) SearchSong(keyword string, opts ...hm) (*apitypes.SearchSongResult, error) {
	opt := hm{}.Merge(opts...).Set("type", apitypes.SearchTypeSong).Set("s", keyword)
	resp, err := api.search(keyword, "cloudsearch", opt)
	if err != nil {
		return nil, err
	}

	var res apitypes.SearchSongResult
	if err = resp.Deserialize(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *NeteaseAPI) SearchAlbum(keyword string, opts ...hm) (*apitypes.SearchAlbumResult, error) {
	opt := hm{}.Merge(opts...).Set("type", apitypes.SearchTypeAlbum).Set("s", keyword)
	resp, err := api.search(keyword, "cloudsearch", opt)
	if err != nil {
		return nil, err
	}
	var res apitypes.SearchAlbumResult
	if err = resp.Deserialize(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *NeteaseAPI) SearchCollection(keyword string, opts ...hm) (*apitypes.SearchCollectionResult, error) {
	opt := hm{}.Merge(opts...).Set("type", apitypes.SearchTypeCollection).Set("s", keyword)
	resp, err := api.search(keyword, "cloudsearch", opt)
	if err != nil {
		return nil, err
	}
	var res apitypes.SearchCollectionResult
	if err = resp.Deserialize(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *NeteaseAPI) SearchArtist(keyword string, opts ...hm) (*apitypes.SearchArtistResult, error) {
	opt := hm{}.Merge(opts...).Set("type", apitypes.SearchTypeArtist).Set("s", keyword)
	resp, err := api.search(keyword, "cloudsearch", opt)
	if err != nil {
		return nil, err
	}
	var res apitypes.SearchArtistResult
	if err = resp.Deserialize(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
