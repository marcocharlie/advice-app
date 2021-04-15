package cache

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
)

func TestAdvicesCache(t *testing.T) {
	var advice []string
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

	//Empty advice
	advicesCache.SetTopicAdvicesCache(topic, advice)

	//Assert advice are empty
	advice, err := advicesCache.GetTopicAdvicesCache(topic)
	if err != nil {
		t.Fatal(err)
	}
	if len(advice) != 0 {
		t.Fatal("Cache is not empty")
	}

	//Insert test data into advice
	advice = []string{
		"advice n. 1",
		"advice n. 22",
		"advice n. 3",
	}

	advicesCache.SetTopicAdvicesCache(topic, advice)

	//Assert correct number of advices is present
	advice, err = advicesCache.GetTopicAdvicesCache(topic)
	if err != nil {
		t.Fatal(err)
	}
	if len(advice) != 3 {
		t.Fatal("Cannot retrieve advice from cache")
	}

}
