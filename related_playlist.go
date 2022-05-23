package ncmapi

// 说明 : 调用此接口,传入歌单 id 可获取相关歌单
//
// required
// 必选参数 : id : 歌单 id
func (api *NeteaseAPI) RelatedPlaylist(id int) (*APIResponse, error) {
	// resp, err := api.Req(http.MethodPost, APIRoutes["related_playlist"], apitypes.H{"id": id}, nil)

	// return _deserializeToImplicitResult(resp, err)
	return nil, nil
}
