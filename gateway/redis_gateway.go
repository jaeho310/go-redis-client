package gateway

type RedisGateway interface {
	SetData(key string, value string) error
	GetData(key string) (string, error)
	GetKeyList() ([]string, error)
}
