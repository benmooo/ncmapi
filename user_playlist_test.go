package ncmapi_test

import "testing"

func TestUserPlaylist(t *testing.T) {
	resp, err := api.UserPlaylist(userID)

	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
	// t.Log(res)
}
