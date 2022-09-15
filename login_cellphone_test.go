package ncmapi_test

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoginPhone(t *testing.T) {
	f, err := ioutil.ReadFile(".auth")
	if err != nil {
		t.Fatal(err)
	}

	auth := strings.Split(string(f), "\n")

	if len(auth) < 2 {
		t.Fatal("not correct auth credential!")
	}
	//
	// print(auth[0], auth[1])

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
