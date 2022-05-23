package apitypes

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	necmcrypto "github.com/benmooo/ncmapi/crypto"
)

type APIRequest struct {
	Method string
	URL    string

	Data H
	// Header http.Header

	Option *RequestOption
}

func NewAPIRequest(method, url string, data H, opt *RequestOption) *APIRequest {
	if data == nil {
		data = H{}
	}
	if opt == nil {
		opt = DefaultRequestOption()
	}
	return &APIRequest{
		Method: method,
		URL:    url,
		Data:   data,
		Option: opt,
	}
}

func (r *APIRequest) ID() string {
	b, _ := json.Marshal(r)
	hash := md5.Sum(b)
	return hex.EncodeToString(hash[:])
}

type RequestOption struct {
	Crypto  necmcrypto.Crypto
	UA      UserAgentType
	Cookies []*http.Cookie
	RealIP  string
	URL     string
	// Proxy url.Proxy,
}

func DefaultRequestOption() *RequestOption {
	return &RequestOption{
		Crypto: necmcrypto.CtryptoWeapi,
		UA:     UAChrome,
	}
}

func (ro *RequestOption) SetURL(u string) *RequestOption {
	ro.URL = u
	return ro
}

func (ro *RequestOption) Cookie(name string) string {
	for _, cookie := range ro.Cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

func (ro *RequestOption) CookieOr(name string, dvalue string) string {
	cookie := ro.Cookie(name)
	if len(cookie) == 0 {
		return dvalue
	}
	return cookie
}

func (ro *RequestOption) AddCookie(name, value string) *RequestOption {
	ro.Cookies = append(ro.Cookies, &http.Cookie{Name: name, Value: value})
	return ro
}

func (ro *RequestOption) AddCookies(cookies ...*http.Cookie) *RequestOption {
	ro.Cookies = append(ro.Cookies, cookies...)
	return ro
}

func (ro *RequestOption) SetCookies(cookies []*http.Cookie) *RequestOption {
	ro.Cookies = cookies
	return ro
}

func (ro *RequestOption) SetUA(ua UserAgentType) *RequestOption {
	ro.UA = ua
	return ro
}

func (ro *RequestOption) SetCrypto(c necmcrypto.Crypto) *RequestOption {
	ro.Crypto = c
	return ro
}

func (ro *RequestOption) SetRealIP(ip string) *RequestOption {
	ro.RealIP = ip
	return ro
}

type UserAgentType uint8

const (
	UAChrome UserAgentType = iota + 1
	UAFirefox
	UAAndroid
	UASafari
	UAIPhone
)
