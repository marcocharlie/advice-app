package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	advisor "github.com/marcocharlie/advice-app/api/internal/advisor"
	"github.com/marcocharlie/advice-app/api/models"
)

// RPCService provides some RPC methods useful to get advice
type RPCService struct {
	redisCLient *redis.Client
}

// NewRPCService instantiates a new RPCService
func NewRPCService(redisClient *redis.Client) *RPCService {
	return &RPCService{
		redisCLient: redisClient,
	}
}

// GiveMeAdvice returns the Advice list matching the given args.Topic and args.Amount, if present, or an error otherwise
func (s *RPCService) GiveMeAdvice(r *http.Request, args *models.AdviceRequestArgs, reply *models.AdviceResponse) error {
	log.Printf("got GiveMeAdvice request from %s\n", r.RemoteAddr)

	advice, err := advisor.GetAdvice(s.redisCLient, args)
	if err != nil {
		log.Println(err)
		return err
	}
	*reply = *advice

	return nil
}
