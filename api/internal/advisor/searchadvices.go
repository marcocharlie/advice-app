package advisor

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	adviceslip "github.com/marcocharlie/advice-app/api/adapter/web"
	cache "github.com/marcocharlie/advice-app/api/data/cache"
	"github.com/marcocharlie/advice-app/api/models"
)

// GetAdvice returns an AdviceResponse, if present, or error otherwise
func GetAdvice(redisClient *redis.Client, adviceRequest *models.AdviceRequestArgs) (*models.AdviceResponse, error) {

	advicesCache := cache.NewAdvicesCache(
		redisClient,
	)

	var adviceMaxAmount = 0
	var adviceList []string
	var adviceResponse models.AdviceResponse
	var existsTopicAdviceCache = true

	err := adviceRequest.Validate()
	if err != nil {
		return nil, err
	}

	// get topic advice from cache, if present
	adviceList, err = advicesCache.GetTopicAdvicesCache(adviceRequest.Topic)
	if err != nil {
		if err == redis.Nil {
			log.Printf("no cache available for the given topic: %s", adviceRequest.Topic)
		} else {
			log.Println(err)
		}
		existsTopicAdviceCache = false
	}

	// if there is not cache for given topic
	if !existsTopicAdviceCache {
		// searches for advice
		advice, err := adviceslip.SearchAdviceSlips(adviceRequest.Topic)
		if err != nil {
			return nil, err
		}

		// lists advice
		for _, advice := range *advice {
			adviceList = append(adviceList, advice.Advice)
		}

		// saves topic advice to local cache
		err = advicesCache.SetTopicAdvicesCache(adviceRequest.Topic, adviceList)
		if err != nil {
			log.Print("setTopicAdvicesCache method errored with: ", err)
		}
	}

	if len(adviceList) == 0 {
		return nil, errors.New(fmt.Sprintf("no advice slips avalable for the given topic: %s", adviceRequest.Topic))
	}

	// set advice max amount if requested
	if adviceRequest.Amount != nil && *adviceRequest.Amount > 0 {
		adviceMaxAmount = *adviceRequest.Amount
	}

	adviceResponse.AdviceList = adviceList
	// limit advice
	if adviceMaxAmount > 0 && len(adviceList) >= adviceMaxAmount {
		adviceResponse.AdviceList = adviceList[:adviceMaxAmount]
	}

	return &adviceResponse, nil
}
