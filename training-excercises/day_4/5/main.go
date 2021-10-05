package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const stringToSearch = "concurrency"

var sites = []string{
	"https://google.com",
	"https://itc.ua/",
	"https://twitter.com/concurrencyinc",
	"https://twitter.com/",
	"http://localhost:8000",
	"https://github.com/bradtraversy/go_restapi/blob/master/main.go",
	"https://www.youtube.com/",
	"https://postman-echo.com/get",
	"https://en.wikipedia.org/wiki/Concurrency_(computer_science)#:~:text=In%20computer%20science%2C%20concurrency%20is,without%20affecting%20the%20final%20outcome.",
}

type SiteData struct {
	data []byte
	uri  string
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	resultsCh := make(chan SiteData, len(sites))

	for _, site := range sites {
		go worker(ctx, resultsCh, site)
	}

	go reader(ctx, cancel, resultsCh)

	time.Sleep(time.Second)
}

func reader(ctx context.Context, cancel context.CancelFunc, r chan SiteData) {
	go func() {
		select {
		case SiteData := <-r:
			if strings.Contains(string(SiteData.data), stringToSearch) {
				fmt.Println("'concurrency' string is found in", SiteData.uri)
				cancel()
			}
		case <-ctx.Done():
		}
	}()
}

func worker(ctx context.Context, r chan SiteData, site string) {
	fmt.Println("starting sending request to", site)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, site, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	r <- SiteData{data: bodyBytes, uri: site}
}
