package necmapi_test

import (
	"testing"
)

func TestLyric(t *testing.T) {
	var songID = 1331628393

	resp, err := api.Lyric(songID)
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
