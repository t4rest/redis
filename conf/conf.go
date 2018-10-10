package conf

type RedisConf struct {
	Address string
}

type HttpConf struct {
	ListenOnPort string
}

type AppConf struct {
	Redis RedisConf
	Http  HttpConf
}

// New
func New() AppConf {
	return AppConf{
		Redis: RedisConf{
			Address: ":12311",
		},
		Http: HttpConf{
			ListenOnPort: ":8080",
		},
	}
}
