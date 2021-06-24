// package store implements the neteaseapi store interface
package inmemstore

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type InMemStore struct {
	Store *cache.Cache

	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

func (s *InMemStore) Add(k string, v interface{}, d time.Duration) error {
	return s.Store.Add(k, v, d)
}

func (s *InMemStore) Set(k string, v interface{}, d time.Duration) {
	s.Store.Set(k, v, d)
}

func (s *InMemStore) Get(k string) (interface{}, bool) {
	return s.Store.Get(k)
}

func (s *InMemStore) Delete(k string) {
	s.Store.Delete(k)
}

func (s *InMemStore) SetDefault(k string, v interface{}) {
	s.Store.SetDefault(k, v)
}

func (s *InMemStore) DeleteExpired() {
	s.Store.DeleteExpired()
}

func (s *InMemStore) Flush() {
	s.Store.Flush()
}

func (s *InMemStore) ItemCount() int {
	return s.Store.ItemCount()
}

func Default() *InMemStore {
	return &InMemStore{
		Store:             cache.New(time.Minute*3, time.Minute*6),
		DefaultExpiration: time.Minute * 3,
		CleanupInterval:   time.Minute * 6,
	}
}

func New(defaultExp, cleanInterval time.Duration) *InMemStore {
	return &InMemStore{
		Store:             cache.New(defaultExp, cleanInterval),
		DefaultExpiration: defaultExp,
		CleanupInterval:   cleanInterval,
	}
}
