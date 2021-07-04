package necmapi_test

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoginPhone(t *testing.T) {
	// persist cookie in local file
	defer api.Client.PreserveCookies()

	f, err := ioutil.ReadFile(".auth")
	if err != nil {
		t.Fatal(err)
	}

	auth := strings.Split(string(f), "\n")

	if len(auth) != 2 {
		t.Fatal("not correct auth credential!")
	}

	resp, err := api.LoginPhone(auth[0], auth[1])
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("got code %d", res.Code)
	}

	t.Logf("%v", "TestLoginPhone passed")
}
