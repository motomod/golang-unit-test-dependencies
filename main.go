package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type httpClient interface {
	get() string
}

type GooglehttpClient struct {
	url string
}

func (client GooglehttpClient) get() string {
	response, err := http.Get(client.url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(responseBody)
}

func main() {
	client := GooglehttpClient{
		"https://google.com",
	}

	fmt.Println(getWebPageFindStringAndReturnLocation(client, "script"))
}

func getWebPageFindStringAndReturnLocation(hc httpClient, find string) int {
	content := hc.get()

	return strings.Index(content, find)
}
