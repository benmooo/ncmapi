package ncmapi_test

import (
	"testing"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

func TestComment(t *testing.T) {
	// resp, err := api.Comment(1, apitypes.ResourceTypeCodeSong, 4262898, apitypes.H{"content": "comments heree"})
	// if err != nil {
	// 	t.Error(err)
	// }

	// res, _ := resp.DeserializeToImplicitResult()
	// if res.Code != 200 {
	// 	t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	// }
	t.Log("Not implemented yet")
}

func TestCommentCreate(t *testing.T) {
	resp, err := api.CommentCreate(songId, apitypes.ResourceTypeCodeSong, "singularity")
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}

func TestCommentDel(t *testing.T) {
	resp, err := api.CommentDel(songId, apitypes.ResourceTypeCodeSong, 5340782540)
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}

func TestCommentRe(t *testing.T) {
	resp, err := api.CommentRe(songId, apitypes.ResourceTypeCodeSong, 5340782540, "pong")
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
