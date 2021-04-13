package cache

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
)

func TestAdvicesCache(t *testing.T) {
	var advices []string
	var topic string

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})

	advicesCache := NewAdvicesCache(
		redisClient,
	)

	topic = "someTopic"

	//Empty advices
	advicesCache.SetTopicAdvicesCache(topic, advices)

	//Assert advices are empty
	advices, err := advicesCache.GetTopicAdvicesCache(topic)
	if err != nil {
		t.Fatal(err)
	}
	if len(advices) != 0 {
		t.Fatal("Cache is not empty")
	}

	//Insert test data into advices
	advices = []string{
		"advice n. 1",
		"advice n. 22",
		"advice n. 3",
	}

	advicesCache.SetTopicAdvicesCache(topic, advices)

	//Assert correct number of advices is present
	advices, err = advicesCache.GetTopicAdvicesCache(topic)
	if err != nil {
		t.Fatal(err)
	}
	if len(advices) != 3 {
		t.Fatal("Cannot retrieve advices from cache")
	}

}
