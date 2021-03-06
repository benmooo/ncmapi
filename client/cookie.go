package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const CookieTmpDir = "/var/tmp/ncmapi_client_cookies"

func (c *Client) Cookies(u *url.URL) []*http.Cookie {
	return c.Jar.Cookies(u)
}

func (c *Client) Cookie(u *url.URL, name string) (*http.Cookie, bool) {
	for _, cookie := range c.Cookies(u) {
		if cookie.Name == name {
			return cookie, true
		}
	}
	return nil, false
}

func (c *Client) HasCookie(u *url.URL, name string) bool {
	for _, cookie := range c.Cookies(u) {
		if cookie.Name == name {
			return true
		}
	}
	return false
}

func (c *Client) WriteCookies(cs []*http.Cookie) error {
	return c.writeCookies(cs)
}

func (c *Client) writeCookies(cs []*http.Cookie) error {
	var rawCookies []string
	for _, c := range cs {
		rawCookies = append(rawCookies, fmt.Sprintf("%s=%s", c.Name, c.Value))
	}

	s := strings.Join(rawCookies, "; ")
	return ioutil.WriteFile(CookieTmpDir, []byte(s), 0644)
}

// func (c *Client) rawCookies() []string {
// 	var rawCookies []string
// 	if c.Config.PreserveCookies {
// 		for _, cookie := range c.Jar.Cookies(&BaseURL) {
// 			rawCookies = append(rawCookies, cookie.String())
// 		}
// 	}
// 	return rawCookies
// }

func (c *Client) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c.Jar.SetCookies(u, cookies)
}

func (c *Client) readLocalCookies() []*http.Cookie {
	b, err := ioutil.ReadFile(CookieTmpDir)
	if err != nil {
		return nil
	}
	rawCookies := strings.Split(string(b), "; ")

	req := http.Request{Header: http.Header{"Cookie": rawCookies}}
	return req.Cookies()
}

func (c *Client) SyncCookiesFromLocal() {
	c.SetCookies(&BaseURL, c.readLocalCookies())
}
