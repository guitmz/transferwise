package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.transferwise.com"
)

type temporaryQuote struct {
	Source           string    `json:"source"`
	Target           string    `json:"target"`
	SourceAmount     float64   `json:"sourceAmount"`
	TargetAmount     float64   `json:"targetAmount"`
	Rate             float64   `json:"rate"`
	CreatedTime      time.Time `json:"createdTime"`
	DeliveryEstimate time.Time `json:"deliveryEstimate"`
	Fee              float64   `json:"fee"`
}

func transferwiseRequest(endpoint string, params map[string]string) []byte {
	u, _ := url.ParseRequestURI(baseURL)
	u.Path = endpoint
	u.RawQuery = mapToQuery(params).Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	check(err)

	res, err := http.DefaultClient.Do(req)
	check(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	check(err)

	return body
}
