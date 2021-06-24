package client

import apitypes "github.com/benmooo/necm-api/api-types"

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
	android: "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	iphone:  "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
}
