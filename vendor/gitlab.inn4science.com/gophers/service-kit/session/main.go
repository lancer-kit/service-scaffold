package session

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type CacheConn interface {
	Init(Settings) error
	GetSessionInfo(string) (SessionInfo, error)
	UpdateSessionInfo(string, SessionInfo) error
}

type Settings struct {
	Host     string
	Port     int
	Password string
	DB       int
}

//TODO: Replace with real struct
type SessionInfo struct {
	SomeField string `json:"some_field"`
}
type RedisCache struct {
	CacheClient *redis.Client
}

func (r *RedisCache) Init(settings Settings) error {
	r.CacheClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", settings.Host, settings.Port),
		Password: settings.Password,
		DB:       settings.DB,
	})

	_, err := r.CacheClient.Ping().Result()
	return err
}

func (r *RedisCache) GetSessionInfo(mark string) (SessionInfo, error) {
	info := SessionInfo{}
	unparsedInfo, err := r.CacheClient.Get(mark).Bytes()
	if err != nil {
		return SessionInfo{}, err
	}

	if err := json.Unmarshal(unparsedInfo, &info); err != nil {
		return SessionInfo{}, err
	}

	return info, nil
}
func (r *RedisCache) UpdateSessionInfo(mark string, info SessionInfo) error {

	ttl, err := r.CacheClient.TTL(mark).Result()
	if err != nil {
		return err
	}

	encodedInfo, err := json.Marshal(info)
	if err := r.CacheClient.Set(mark, encodedInfo, ttl).Err(); err != nil {
		return err
	}
	return nil
}
