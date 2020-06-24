package baidu

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	response, err := GetAccessToken("appID", "secret")
	if err != nil {
		t.Fatal(err)
	}

	println(response.ErrCode)
	println(response.ErrMsg)
}