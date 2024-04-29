package url

import (
	"errors"
	"fmt"
	"net/http"
)

var requestFailed = "Request failed"
var errRequestFailed = errors.New(requestFailed)

func HitURL(urls []string) {
	var results = make(map[string]string)
	
	for _, url := range urls {
		result := "OK"
		err := hit(url)
		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hit(url string) error{
	resp, err := http.Get(url)
	if err != nil ||  resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}