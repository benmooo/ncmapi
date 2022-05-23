package ncmapi_test

import (
	"testing"
)

func TestSearch(t *testing.T) {
	resp, err := api.Search("123")
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}

func TestCloudSearch(t *testing.T) {
	resp, err := api.CloudSearch("xiatian ljj", hm{"limit": 1})
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
func TestSearchSong(t *testing.T) {
	res, err := api.SearchSong("mota")
	if err != nil {
		t.Error(err)
	}

	if res.Code != 200 {
		t.Errorf("code: %d", res.Code)
	}
}

func TestSearchAlbum(t *testing.T) {
	res, err := api.SearchAlbum("出山", hm{"limit": 1})
	if err != nil {
		t.Error(err)
	}

	if res.Code != 200 {
		t.Errorf("code: %d", res.Code)
	}
}

func TestSearchCollection(t *testing.T) {
	res, err := api.SearchCollection("出山", hm{"limit": 1})
	if err != nil {
		t.Error(err)
	}

	if res.Code != 200 {
		t.Errorf("code: %d", res.Code)
	}
}

func TestSearchArtist(t *testing.T) {
	res, err := api.SearchArtist("xusong", hm{"limit": 1})
	if err != nil {
		t.Error(err)
	}

	if res.Code != 200 {
		t.Errorf("code: %d", res.Code)
	}
	// t.Log(res)
}
