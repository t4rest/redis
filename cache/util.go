package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Get get value by key
func (rds *Redis) Get(key string) ([]byte, error) {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}

	return data, err
}

// Set set value by key
func (rds *Redis) Set(key string, value []byte) error {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

// Exists check if value exists
func (rds *Redis) Exists(key string) (bool, error) {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

// Delete delete value by key
func (rds *Redis) Delete(key string) error {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck

	_, err := conn.Do("DEL", key)
	return err
}

// GetKeys ket keys
func (rds *Redis) GetKeys(pattern string) ([]string, error) {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck, gosec

	iter := 0
	var keys []string
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)   // nolint: errcheck, gosec
		k, _ := redis.Strings(arr[1], nil) // nolint: errcheck, gosec
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

// Incr increment counter
func (rds *Redis) Incr(counterKey string) (int, error) {

	conn := rds.pool.Get()
	defer conn.Close() // nolint: errcheck

	return redis.Int(conn.Do("INCR", counterKey))
}
