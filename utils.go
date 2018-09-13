package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func mapToQuery(params map[string]string) url.Values {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	return values
}
