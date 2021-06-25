package necmapi_test

import (
	"net/http"
	"testing"
	"time"

	necmapi "github.com/benmooo/necm-api"
	apitypes "github.com/benmooo/necm-api/api-types"
)

// test request & cache mechnism
func TestApi(t *testing.T) {
	payload := hm{"type": apitypes.SearchTypeSong, "offset": 0, "limit": 1}.Set("s", "xusong")
	opt := apitypes.DefaultRequestOption()
	areq := apitypes.NewAPIRequest(http.MethodPost, necmapi.APIRoutes["cloudsearch"], payload, opt)

	// flush all cache
	api.Store.Flush()

	ids := []string{}
	// request every 10 seconds
	for i := range [5]byte{} {
		t.Logf("Round %d:--------------------------", i+1)

		id := api.Client.RequestID(areq)
		t.Logf("Request ID: %v", id)
		if !contains(ids, id) {
			ids = append(ids, id)
		}

		resp, err := api.Request(areq)
		if err != nil {
			t.Fatal(err)
		}

		res, err := resp.DeserializeToImplicitResult()
		if err != nil {
			t.Fatal(err)
		}

		if res.Code != 200 {
			t.Fatalf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
		}

		time.Sleep(time.Second * 5)
	}

	// cache repository test
	if api.Store.ItemCount() != len(ids) {
		t.Errorf("http requests: %v, got %v response cached", len(ids), api.Store.ItemCount())
	}
	for _, id := range ids {
		if !api.HasCache(id) {
			t.Errorf("response data not cached with request id: %v", id)
		}
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

//
//
//
// Setup

var api = necmapi.New(
	&necmapi.NeteaseAPIConfig{
		CacheDefaultExpiration: time.Minute * 1,
		PreserveCookies:        true,
		// LogHttpResponse:        true,
		LogHttpRequest: true,
	},
)
