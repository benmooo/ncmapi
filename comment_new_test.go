package ncmapi_test

import (
	"testing"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

func TestCommentNew(t *testing.T) {
	resp, err := api.CommentNew(songId, apitypes.ResourceTypeCodeSong, hm{"pageSize": 1})
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}

func TestGetComment(t *testing.T) {
	resp, err := api.GetComment(songId, apitypes.ResourceTypeCodeSong, 1, 1, 1, 0, true)
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
