package necmapi_test

import (
	"testing"
)

func TestLoginPhone(t *testing.T) {
	// persist cookie in local file
	defer api.Client.PreserveCookies()

	phone, password := "13588244781", "vaneyue0802"
	resp, err := api.LoginPhone(phone, password)
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("got code %d", res.Code)
	}

	t.Logf("%v", "TestLoginPhone passed")
}
