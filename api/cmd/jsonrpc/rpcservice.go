package main

import (
	"github.com/marcocharlie/advice-app/api/models"
)

// RPCService provides some RPC methods useful to get advices
type RPCService struct{}

// NewRPCService instantiates a new RPCService
func NewRPCService() *RPCService {
	return &RPCService{}
}

// GiveMeAdvice returns the Advice list matching the given args.Topic and args.Amount, if present, or an error otherwise
func (s *RPCService) GiveMeAdvice(args *models.AdviceRequestArgs) error {

	return nil
}
