package ncmapi_test

import (
	"testing"

	apitypes "github.com/benmooo/ncmapi/api-types"
)

type hm = apitypes.H

func TestScrobble(t *testing.T) {
	songID := 29774784 // Domestic pressure
	resp, err := api.Scrobble(songID, collectionID, hm{"time": 211})
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
