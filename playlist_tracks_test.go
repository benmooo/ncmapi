package ncmapi_test

import "testing"

// my necmID
var userID = 49668844

// playlistID
var collectionID = 2484967117
var songId = 29106885 // An ending, a beginning

func TestPlaylistTracks(t *testing.T) {
	resp, err := api.PlaylistTracks("add", collectionID, []int{songId})
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
