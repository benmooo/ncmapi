package ncmapi_test

import "testing"

func TestArtistSongs(t *testing.T) {
	resp, err := api.ArtistSongs(99392)
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
