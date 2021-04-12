package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func main() {

	// instantiate a new RPC Server and register appropriate codecs
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	// create a new RPCService and register it to the RPC Server
	rpcService := NewRPCService()
	s.RegisterService(rpcService, "RPCService")

	// prepare an HTTP Router, pointing the /rpc url to the RPC Server
	r := mux.NewRouter()
	r.Handle("/rpc", s)

	// start HTTP server on port 5000
	log.Fatal(http.ListenAndServe(":5000", r))
}