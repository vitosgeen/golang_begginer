package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Posts []struct {
	Id int `json:"id"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var cntPosts = 0

func main() {
	postsByte := getExternalDataByteByUrl("https://jsonplaceholder.typicode.com/posts")
	posts := &Posts{}
	errUnM := json.Unmarshal(postsByte, &posts)
	if errUnM != nil {
		log.Println(errUnM)
	}
	for post := range *posts {
		go getPostById(post)
	}
	var input string
	fmt.Scanln(&input)
}

func getPostById(id int) {
	url := "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(id)
	postByte := getExternalDataByteByUrl(url)
	post := &Post{}
	errUnM := json.Unmarshal(postByte, &post)
	if errUnM != nil {
		log.Println(errUnM)
	}
	if cntPosts < 4 {
		fmt.Println(post)
	}
	cntPosts++
}

func getExternalDataByteByUrl(url string) []byte {
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
	return body
}
