package cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/t4rest/redis/conf"
)

// Redis redis struct
type Redis struct {
	pool *redis.Pool
}

// New create new pool redis connections
func New(cfg conf.AppConf) (CacheInterface, error) {
	rds := &Redis{
		pool: newPool(cfg.Redis.Address),
	}
	err := rds.ping()

	return rds, err
}

// Close pool of connections
func (rds *Redis) Close() {
	rds.pool.Close() // nolint: errcheck, gosec
}

func (rds *Redis) ping() error {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, nil
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
