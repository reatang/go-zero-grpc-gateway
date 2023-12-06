package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/reatang/go-zero-grpc-gateway/example/simple_rpc/pb/simple_rpc"
)

var host = "http://127.0.0.1:8888"

func main() {
	req := simple_rpc.Request{Ping: "123"}
	body, _ := json.Marshal(req)

	response, err := http.Post(host+"/grpc1/simple/ping", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", respBody)
}
