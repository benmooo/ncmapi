package necmapi

import (
	"fmt"
	"net/http"
	"strings"

	apitypes "github.com/benmooo/necm-api/api-types"
)

// 说明 : 登录后调用此接口,可以更新用户歌单
//
// required
// id:歌单id
// name:歌单名字
// desc:歌单描述
// tags:歌单tag ,多个用 `;` 隔开,只能用官方规定标签
func (api *NeteaseAPI) PlaylistUpdate(id int, name, desc string, tags []string) (*APIResponse, error) {
	opt := apitypes.DefaultRequestOption().AddCookie("os", "pc")
	data := playlistUpdatePayload(id, name, desc, tags)

	resp, err := api.Req(http.MethodPost, APIRoutes["playlist_update"], data, opt)
	return resp, err
}

func playlistUpdatePayload(id int, name, desc string, tags []string) hm {
	return hm{
		"/api/playlist/update/name": fmt.Sprintf(`{"id":%d,"name":"%s"}`, id, name),
		"/api/playlist/desc/update": fmt.Sprintf(`{"id":%d,"desc":"%s"}`, id, desc),
		"/api/playlist/tags/update": fmt.Sprintf(`{"id":%d,"tags":"%s"}`, id, strings.Join(tags, ";")),
	}
}
