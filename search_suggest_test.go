package ncmapi_test

import "testing"

func TestSearchSuggest(t *testing.T) {
	resp, err := api.SearchSuggest("xus")

	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
