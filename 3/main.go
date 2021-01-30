package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	getExternalDataJson(url)
}

func getExternalDataJson(url string) {
	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	// get []byte data from request
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
