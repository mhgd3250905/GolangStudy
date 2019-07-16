package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	httpUrl := "http://49.234.76.105/ping"

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://222.76.75.7:5781")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	resp, err := client.Get(httpUrl)
	if err != nil {
		fmt.Println("error : ", err)
		return
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}


}