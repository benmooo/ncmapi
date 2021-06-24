package necmapi_test

import (
	"testing"
)

func TestLoginStatus(t *testing.T) {
	resp, err := api.LoginStatus()
	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
