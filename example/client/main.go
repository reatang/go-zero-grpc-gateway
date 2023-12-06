package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

var host = "http://127.0.0.1:8888"

func main() {
	body := []byte(`{"ping":"123"}`)

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
