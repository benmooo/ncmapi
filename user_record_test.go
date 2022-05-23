package ncmapi_test

import "testing"

func TestUserRecord(t *testing.T) {
	resp, err := api.UserRecord(userID)
	// t.Log(string(resp.Data))

	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
