package ncmapi

import (
	"fmt"
	"net/http"
)

// 说明: 调用此接口,传入歌曲 id, 可获取音乐是否可用,返回 { success: true, message: 'ok' } 或者 { success: false, message: '亲爱的,暂无版权' }
// requried
// 必选参数 : id : 歌曲 id
// optional
// 可选参数 : br: 码率,默认设置了 999000 即最大码率,如果要 320k 则可设置为 320000,其他类推
func (api *NeteaseAPI) CheckMusic(id int, opts ...hm) (*APIResponse, error) {
	data := hm{"br": 999000}.Merge(opts...).Merge(hm{"ids": fmt.Sprintf("[%d]", id)})

	resp, err := api.Req(http.MethodPost, APIRoutes["check_music"], data, nil)
	return resp, err
}
