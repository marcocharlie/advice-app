package cache

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"github.com/marcocharlie/advice-app/api/config"
)

//AdvicesMemory implements the Advice interface using local in memory persistance
type AdvicesCache struct {
	redisClient *redis.Client
}

func NewAdvicesCache(redisClient *redis.Client) AdvicesCache {
	return AdvicesCache{
		redisClient: redisClient,
	}
}

func (ac *AdvicesCache) GetTopicAdvicesCache(topic string) ([]string, error) {

	var advicesCache []string

	topicAdviceCache, err := ac.redisClient.Get(topic).Result()
	if err != nil {
		return advicesCache, err
	}

	err = json.Unmarshal([]byte(topicAdviceCache), &advicesCache)
	if err != nil {
		log.Println(err)
		return advicesCache, err
	}

	log.Printf("retrieved local cache for topic %s", topic)

	return advicesCache, nil

}

func (ac *AdvicesCache) SetTopicAdvicesCache(topic string, advice []string) error {

	json, err := json.Marshal(advice)
	if err != nil {
		return err
	}

	_, err = ac.redisClient.Set(topic, json, config.CacheExpirationTime).Result()
	if err != nil {
		return err
	}

	log.Printf("created local cache for topic %s", topic)

	return nil
}
