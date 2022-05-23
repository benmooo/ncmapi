package client

import apitypes "github.com/benmooo/ncmapi/api-types"

func FakeUA(ua apitypes.UserAgentType) string {
	switch ua {
	case apitypes.UAChrome:
		return useragent.chrome
	case apitypes.UAFirefox:
		return useragent.firefox
	case apitypes.UASafari:
		return useragent.safari
	case apitypes.UAAndroid:
		return useragent.android
	case apitypes.UAIPhone:
		return useragent.iphone

	default:
		return useragent.chrome
	}

}

var useragent = struct {
	chrome  string
	firefox string
	safari  string
	android string
	iphone  string
}{
	chrome:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	firefox: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	safari:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	android: "Mozilla/5.0 (Linux; Android 9; PCT-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.311 Mobile Safari/537.36",
	iphone:  "Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
}
