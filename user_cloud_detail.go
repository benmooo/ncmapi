package ncmapi

import (
	"net/http"
)

// 说明 : 登录后调用此接口 , 传入云盘歌曲 id，可获取云盘数据详情

// requried
// 必选参数 : id: 歌曲id,可多个,用逗号隔开
func (api *NeteaseAPI) UserCloudDetail(ids []int) (*APIResponse, error) {
	resp, err := api.Req(http.MethodPost, APIRoutes["user_cloud_detail"], hm{"songIds": ids}, nil)

	return resp, err
}
