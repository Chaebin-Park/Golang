package url

import (
	"errors"
	"fmt"
	"net/http"
)

var requestFailed = "Request failed"
var errRequestFailed = errors.New(requestFailed)

type result struct {
	url string
	status string
}

func GoRoutineHitURL(urls []string) {
	var results = make(map[string]string)
	c := make(chan result)

	for _, url := range urls {
		go goHitURL(url, c)
	}

	for i:=0;i<len(urls);i++ {
		r := <-c
		results[r.url] = r.status
	}

	for url, status := range(results) {
		fmt.Println(url, "\t", status)
	}
}	

func goHitURL(url string, c chan<- result) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil ||  resp.StatusCode >= 400{
		status = "Failed"
	}
	c <- result{url: url, status: status}
}

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