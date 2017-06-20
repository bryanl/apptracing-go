package main

import (
	"io/ioutil"
	"net/http"

	"github.com/bryanl/apptracing-go/internal/platform/logging"
)

func main() {
	logger := logging.Init("client")

	req, _ := http.NewRequest("GET", "http://localhost:8081", nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Log("msg", "http client error", "err", err.Error())
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	logger.Log("msg", string(b))
}
