package service

import "github.com/t4rest/redis/cache"

type (
	// Worker .
	Worker interface {
		Work() string
	}

	// service implements the Fetcher.
	service struct {
		cache cache.CacheInterface
	}
)

// New .
func New(cache cache.CacheInterface) Worker {
	return &service{cache: cache}
}

// Work .
func (srv *service) Work() string {

	if exist, _ := srv.cache.Exists("work"); exist {
		data, err := srv.cache.Get("work")
		if err != nil {
			return ""
		}
		return string(data)
	}

	srv.cache.Set("work", []byte("work"))

	return ""
}
