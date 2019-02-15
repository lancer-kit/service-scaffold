package session

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gitlab.inn4science.com/gophers/service-kit/log"

	"github.com/stretchr/testify/require"

	"github.com/go-redis/redis"
)

func TestSession(t *testing.T) {
	//init redis with some test records
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	require.NoErrorf(t, err, "Unable to establish conn with Redis")

	for i := 1; i <= 5; i++ {
		encodedInfo, err := json.Marshal(SessionInfo{SomeField: "test info"})
		require.NoError(t, err)

		err = client.Set(fmt.Sprintf("mark %d", i),
			encodedInfo,
			time.Hour,
		).Err()
		require.NoErrorf(t, err, "Unable to write into redis")
	}

	//init client
	cacheClient := RedisCache{}
	cacheClient.Init(Settings{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	})

	//verify GetSessionInfo
	for i := 1; i <= 5; i++ {
		info, err := cacheClient.GetSessionInfo(fmt.Sprintf("mark %d", i))
		require.NoErrorf(t, err, "Unable to read from redis")

		log.Default.WithField("obtained info:", info).Info("TEST session client")
	}

	//verify UpdateSessionInfo
	for i := 1; i <= 5; i++ {
		err := cacheClient.UpdateSessionInfo(fmt.Sprintf("mark %d", i), SessionInfo{SomeField: "updated test info"})
		require.NoErrorf(t, err, "Unable to update data in redis")
	}

	//verify GetSessionInfo
	for i := 1; i <= 5; i++ {
		info, err := cacheClient.GetSessionInfo(fmt.Sprintf("mark %d", i))
		require.NoErrorf(t, err, "Unable to read from redis")

		log.Default.WithField("obtained info:", info).Info("TEST session client")
	}
}
