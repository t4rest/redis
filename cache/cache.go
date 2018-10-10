package cache

// CacheInterface base cache interface
type CacheInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Exists(key string) (bool, error)
	Delete(key string) error
	Close()
}
