// this package is a basic implementation of APIClient
package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	apitypes "github.com/benmooo/ncmapi/api-types"
	necmcrypto "github.com/benmooo/ncmapi/crypto"
)

var BaseURL = url.URL{
	Scheme: "https",
	Host:   "music.163.com",
}

type Client struct {
	*http.Client
	Config *ClientConfig
}

type ClientConfig struct {
	// if true, client will read cookies from local & preserve cookie to local when exit
	PreserveCookies bool

	// logs
	LogHttpRequest  bool
	LogHttpResponse bool
}

func NewClient(cfg *ClientConfig) *Client {
	c := &Client{
		&http.Client{
			Jar: defaultCookieJar(),
		},
		cfg,
	}
	c.init()
	return c
}

func (c *Client) init() {
	if c.Config.PreserveCookies {
		c.SyncCookiesFromLocal()
	}
}

func Default() *Client {

	c := &Client{
		&http.Client{
			Jar: defaultCookieJar(),
		},
		DefaultConfig(),
	}
	c.init()
	return c
}

func DefaultConfig() *ClientConfig {
	return &ClientConfig{
		PreserveCookies: true,
		LogHttpRequest:  false,
		LogHttpResponse: false,
	}
}

func (c *Client) Req(method string, url string, data apitypes.H, opt *apitypes.RequestOption) (*apitypes.APIResponse, error) {
	return c.req(apitypes.NewAPIRequest(method, url, data, opt))
}

func (c *Client) Request(r *apitypes.APIRequest) (*apitypes.APIResponse, error) {
	return c.req(r)
}

func (c *Client) req(r *apitypes.APIRequest) (*apitypes.APIResponse, error) {
	req, err := c.toHttpRequest(r)
	if err != nil {
		return nil, err
	}

	if c.Config.LogHttpRequest {
		log.Printf("http request: %+v", req)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	if c.Config.LogHttpResponse {
		log.Printf("http response cookies: %+v", resp.Cookies())
		log.Printf("http response: %+v", string(data))
	}
	// This is a compromise way to figure out that the reponse body are encryped or not
	//
	// According to https://github.com/Binaryify/NeteaseCloudMusicApi/util/request.js line: 160
	// of version: 122b6f765b229ede1b04844f99945271217dd732
	//
	// If there is a better way, please let me known
	if r.Option.Crypto == necmcrypto.CtryptoEapi {
		// try deserialization
		if err = json.Unmarshal(data, &hm{}); err != nil {
			pt, err := necmcrypto.EapiDecrypt(data)
			// can't deserialization or decrypt --> might be bad request
			// something like this: {code: 400, msg: bad request}{code: 400, msg: bad request}
			if err != nil {
				return nil, fmt.Errorf("%v with: %v", err.Error(), string(data))
			}
			return &apitypes.APIResponse{Data: pt}, nil
		}
	}

	// sync cookies
	cs := resp.Cookies()
	if c.Config.PreserveCookies && len(cs) > 0 {
		c.WriteCookies(cs)
	}

	return &apitypes.APIResponse{Data: data}, nil
}

func (c *Client) RequestID(r *apitypes.APIRequest) string {
	return c.decorateAPIRequest(r).ID()
}

func (c *Client) ReqID(method string, url string, data apitypes.H, opt *apitypes.RequestOption) string {
	return c.RequestID(apitypes.NewAPIRequest(method, url, data, opt))
}

func (c *Client) toHttpRequest(r *apitypes.APIRequest) (*http.Request, error) {
	if r.Option.Crypto == necmcrypto.CtryptoLinuxapi {
		return nil, &ClientErr{Msg: "Crypto 4 Linuxapi not implemented yet!"}
	}
	r = c.decorateAPIRequest(r)

	payload := strings.NewReader(c.values(r).Encode())
	req, _ := http.NewRequest(r.Method, r.URL, payload)
	req.Header = c.header(r)
	return req, nil
}

func (c *Client) header(r *apitypes.APIRequest) http.Header {
	header := http.Header{}
	header.Add("User-Agent", FakeUA(r.Option.UA))
	// header.Add("X-Real-IP", "118.88.88.88")
	if r.Method == http.MethodPost {
		header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if strings.Contains(r.URL, BaseURL.Hostname()) {
		header.Add("Referer", BaseURL.String())
	}
	if len(r.Option.RealIP) > 0 {
		header.Set("X-Real-IP", r.Option.RealIP)
	}
	for _, cookie := range r.Option.Cookies {
		header.Add("Cookie", cookie.String())
	}

	switch r.Option.Crypto {
	case necmcrypto.CtryptoWeapi:
	case necmcrypto.CtryptoLinuxapi:
	case necmcrypto.CtryptoEapi:
		for _, cookie := range c.cookies(r) {
			if len(cookie.Value) > 0 {
				header.Add("Cookie", cookie.String())
			}
		}
	}

	return header
}

func (c *Client) values(r *apitypes.APIRequest) url.Values {
	data := url.Values{}

	switch r.Option.Crypto {
	case necmcrypto.CtryptoWeapi:
		encrypted := necmcrypto.WeapiEncrypt(r.Data.Json())
		for k, v := range encrypted {
			data.Add(k, v)
		}
	case necmcrypto.CtryptoLinuxapi:
		// Todo
	case necmcrypto.CtryptoEapi:
		encrypted := necmcrypto.EapiEncrypt([]byte(r.Option.URL), r.Data)
		for k, v := range encrypted {
			data.Add(k, v)
		}
	}

	return data
}

func (c *Client) decorateAPIRequest(r *apitypes.APIRequest) *apitypes.APIRequest {
	r = c.decoUrl(r)
	// data
	r = c.decoData(r)
	// options & cookies
	r = c.decoCookies(r)
	return r
}

func (c *Client) decoUrl(r *apitypes.APIRequest) *apitypes.APIRequest {
	re := regexp.MustCompile(`\w*api`)

	switch r.Option.Crypto {
	case necmcrypto.CtryptoWeapi:
		r.URL = re.ReplaceAllString(r.URL, "weapi")
	case necmcrypto.CtryptoLinuxapi:
		// ToDo
	case necmcrypto.CtryptoEapi:
		r.URL = re.ReplaceAllString(r.URL, "eapi")
	}

	return r
}

func (c *Client) decoData(r *apitypes.APIRequest) *apitypes.APIRequest {
	switch r.Option.Crypto {
	case necmcrypto.CtryptoWeapi:
		r.Data["csrf_token"] = ""
		for _, cookie := range c.Client.Jar.Cookies(&BaseURL) {
			if cookie.Name == "__csrf" && len(cookie.Value) > 0 {
				r.Data["csrf_token"] = cookie.Value
				break
			}
		}
	case necmcrypto.CtryptoLinuxapi:
		// ToDo
	case necmcrypto.CtryptoEapi:
		header := hm{}
		for _, cookie := range c.cookies(r) {
			header.Set(cookie.Name, cookie.Value)
		}
		r.Data.Set("header", header)
	}

	return r
}

// The client cookie jar will persist the cookies responsed by the server
// and following request will take along these cookies which is the default behavior
// defined in http.Client

// However, we still need present RequestOption with cookie which resides api.Client.Jar
// The reason is: The second APIRequest considered different from the the first one if
// cookies updated by server even if two APIRequest is identical
func (c *Client) decoCookies(r *apitypes.APIRequest) *apitypes.APIRequest {
	u, _ := url.Parse(r.URL)

	// present request with cookies
	for _, cookie := range c.Client.Jar.Cookies(u) {
		if !contains(r.Option.Cookies, cookie) {
			r.Option.AddCookies(cookie)
		}
	}
	return r
}

func (c *Client) cookies(r *apitypes.APIRequest) []*http.Cookie {
	var cookies []*http.Cookie
	// request cookie
	cs := map[string]string{
		"osver":       r.Option.Cookie("osver"),
		"deviceId":    r.Option.Cookie("deviceId"),
		"appver":      r.Option.CookieOr("appver", "8.0.0"),
		"versioncode": r.Option.CookieOr("versioncode", "140"),
		"mobilename":  r.Option.Cookie("mobilename"),
		"buildver":    r.Option.CookieOr("buildver", strconv.Itoa(int(time.Now().Unix()))),
		"resolution":  r.Option.CookieOr("resolution", "1920x1080"),
		"__csrf":      r.Option.Cookie("__csrf"),
		"os":          r.Option.CookieOr("os", "android"),
		"channel":     r.Option.Cookie("channel"),
		"requestId":   fmt.Sprintf("%d_%04d", time.Now().Unix(), rand.Intn(1000)),
		"MUSIC_U":     r.Option.Cookie("MUSIC_U"),
		"MUSIC_A":     r.Option.Cookie("MUSIC_A"),
	}
	for k, v := range cs {
		cookies = append(cookies, &http.Cookie{Name: k, Value: v})
	}

	return cookies
}

func defaultCookieJar() *cookiejar.Jar {
	jar, _ := cookiejar.New(nil)
	return jar
}

func contains(cs []*http.Cookie, c *http.Cookie) bool {
	for _, cookie := range cs {
		if cookie.Name == c.Name {
			return true
		}
	}
	return false
}

type hm = apitypes.H
