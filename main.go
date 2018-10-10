package main

import (
	"github.com/t4rest/redis/api"
	"github.com/t4rest/redis/cache"
	"github.com/t4rest/redis/conf"
	"github.com/t4rest/redis/service"
)

func main() {

	cfg := conf.New()

	c, err := cache.New(cfg)
	if err != nil {
		panic(err)
	}
	srv := service.New(c)

	apiModule := api.New(cfg, srv)
	apiModule.StartServe()
}
