package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func main() {

	// init redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})

	// instantiate a new RPC Server and register appropriate codecs
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	// create a new RPCService and register it to the RPC Server
	rpcService := NewRPCService(redisClient)
	s.RegisterService(rpcService, "RPCService")

	// prepare an HTTP Router, pointing the /api/advice url to the RPC Server
	r := mux.NewRouter()
	r.Handle("/api/advice", s)

	// enable cors
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST"})

	// start HTTP server on port 5000
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
