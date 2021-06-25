package necmapi

import (
	"net/http"
	"net/url"
	"time"

	apitypes "github.com/benmooo/necm-api/api-types"
	"github.com/benmooo/necm-api/client"
	"github.com/benmooo/necm-api/inmemstore"
)

var BaseURL = url.URL{
	Scheme: "https",
	Host:   "music.163.com",
}

type NeteaseAPI struct {
	Client APIClient
	Store  Store
	Config *NeteaseAPIConfig
}

func New(cfg *NeteaseAPIConfig) *NeteaseAPI {
	api := &NeteaseAPI{
		Client: client.NewClient(
			&client.ClientConfig{
				PreserveCookies: cfg.PreserveCookies,
				LogHttpRequest:  cfg.LogHttpRequest,
				LogHttpResponse: cfg.LogHttpResponse,
			}),
		Config: cfg,
	}
	if !cfg.DisableCache {
		if cfg.CacheDefaultExpiration == 0 {
			cfg.CacheDefaultExpiration = time.Minute * 3
		}
		if cfg.CacheCleanupInterval == 0 {
			cfg.CacheCleanupInterval = time.Minute * 6
		}
		api.Store = inmemstore.New(
			cfg.CacheDefaultExpiration, cfg.CacheCleanupInterval)
	}

	return api
}

func Default() *NeteaseAPI {
	cfg := DefaultNeteaseAPIConfig()
	return New(cfg)
}

// This interface defines functions that a client model possess
type APIClient interface {
	Request(areq *apitypes.APIRequest) (*apitypes.APIResponse, error)
	Req(method string, url string, data apitypes.H, opt *apitypes.RequestOption) (*apitypes.APIResponse, error)
	RequestID(areq *apitypes.APIRequest) string
	ReqID(method string, url string, data apitypes.H, opt *apitypes.RequestOption) string

	// cookies synchronization
	ReadSetCookies()
	PreserveCookies() error
	Cookies(u *url.URL) []*http.Cookie
	HasCookie(u *url.URL, name string) bool
}

type Store interface {
	// Add an item to the cache only if an item doesn't already exist, or if the existing item has expired. Returns an error otherwise.
	Add(k string, v interface{}, d time.Duration) error
	// Delete an item from the cache. Does nothing if the key is not in the cache.
	Delete(k string)
	// Delete all expired items from the cache.
	DeleteExpired()
	// Delete all items from the cache.
	Flush()
	//Get an item from the cache
	Get(k string) (interface{}, bool)
	//Add an item to the cache, replacing any existing item. If the duration is 0 (DefaultExpiration), the cache's default expiration time is used.
	Set(k string, v interface{}, d time.Duration)
	SetDefault(k string, v interface{})
	// Returns the number of items in the cache. This may include items that have expired, but have not yet been cleaned up.
	ItemCount() int
}

func (api *NeteaseAPI) Req(method string, url string, data apitypes.H, opt *apitypes.RequestOption) (*apitypes.APIResponse, error) {
	req := apitypes.NewAPIRequest(method, url, data, opt)
	return api.Request(req)
}

func (api *NeteaseAPI) Request(areq *apitypes.APIRequest) (*apitypes.APIResponse, error) {
	return api.resolveRequest(areq)
}

func (api *NeteaseAPI) resolveRequest(areq *apitypes.APIRequest) (*apitypes.APIResponse, error) {
	id := api.Client.RequestID(areq)

	// has cache?
	if !api.Config.DisableCache {
		if data, ok := api.readFromCache(id); ok {
			return data, nil
		}
	}

	resp, err := api.Client.Request(areq)
	if err != nil {
		return nil, err
	}

	// cache response
	if !api.Config.DisableCache {
		api.cacheAPIResponseDefault(id, resp)
	}
	return resp, nil
}

func (api *NeteaseAPI) HasCache(id string) bool {
	_, ok := api.Store.Get(id)
	return ok
}

func (api *NeteaseAPI) readFromCache(id string) (*apitypes.APIResponse, bool) {
	res, ok := api.Store.Get(id)
	if !ok {
		return nil, false
	}
	return res.(*apitypes.APIResponse), true
}

func (api *NeteaseAPI) cacheAPIResponse(id string, resp *apitypes.APIResponse, d time.Duration) {
	api.Store.Set(id, resp, d)
}

func (api *NeteaseAPI) cacheAPIResponseDefault(id string, resp *apitypes.APIResponse) {
	api.cacheAPIResponse(id, resp, 0)
}

type NeteaseAPIConfig struct {
	DisableCache           bool
	CacheDefaultExpiration time.Duration
	CacheCleanupInterval   time.Duration

	PreserveCookies bool
	LogHttpRequest  bool
	LogHttpResponse bool
}

func DefaultNeteaseAPIConfig() *NeteaseAPIConfig {
	return &NeteaseAPIConfig{
		CacheDefaultExpiration: time.Minute * 3,
		CacheCleanupInterval:   time.Minute * 6,
		PreserveCookies:        true,
	}

}

type hm = apitypes.H
type APIResponse = apitypes.APIResponse
