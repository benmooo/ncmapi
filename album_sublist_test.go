package necmapi_test

import "testing"

func TestAlbumSublist(t *testing.T) {
	resp, err := api.AlbumSublist()
	if err != nil {
		t.Error(err)
	}

	res, err := resp.DeserializeToImplicitResult()
	if err != nil {
		t.Fatal(err)
	}
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
