package necmapi_test

import "testing"

func TestUserCloudDetail(t *testing.T) {
	ids := []int{1853497562, 1485489697}
	resp, err := api.UserCloudDetail(ids)

	if err != nil {
		t.Error(err)
	}

	res, _ := resp.DeserializeToImplicitResult()
	if res.Code != 200 {
		t.Errorf("code: %d, msg: %s, message: %s, time: %d", res.Code, res.Msg, res.Message, res.Time)
	}
}
