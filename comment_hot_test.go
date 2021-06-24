package necmapi_test

import (
	"testing"

	apitypes "github.com/benmooo/necm-api/api-types"
)

func TestCommentHot(t *testing.T) {
	resp, err := api.CommentHot(songId, apitypes.ResourceTypeCodeSong)
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
