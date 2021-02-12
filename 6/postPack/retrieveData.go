package postPack

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const urlComments = "https://jsonplaceholder.typicode.com/comments"
const urlPosts = "https://jsonplaceholder.typicode.com/posts"

func GetCommentsByPostId(id int) *[]Comment {
	url := urlComments + "?postId=" + strconv.Itoa(id)
	commentsByte := getExternalDataByteByUrl(url)
	comments := &[]Comment{}
	errUnM := json.Unmarshal(commentsByte, &comments)
	if errUnM != nil {
		log.Println(errUnM)
	}
	return comments
}
func GetPostsByUserId(id int) *[]Post {
	url := urlPosts + "?userId=" + strconv.Itoa(id)
	postsByte := getExternalDataByteByUrl(url)
	posts := &[]Post{}
	errUnM := json.Unmarshal(postsByte, &posts)
	if errUnM != nil {
		log.Println(errUnM)
	}
	return posts
}
func GetPostsIdByUserId(id int) *Posts {
	url := urlPosts + "?userId=" + strconv.Itoa(id)
	postsByte := getExternalDataByteByUrl(url)
	posts := &Posts{}
	errUnM := json.Unmarshal(postsByte, &posts)
	if errUnM != nil {
		log.Println(errUnM)
	}
	return posts
}
func GetPostsId() *Posts {
	postsByte := getExternalDataByteByUrl(urlPosts)
	posts := &Posts{}
	errUnM := json.Unmarshal(postsByte, &posts)
	if errUnM != nil {
		log.Println(errUnM)
	}
	return posts
}

func GetPostById(id int) *Post {
	url := urlPosts + "/" + strconv.Itoa(id)
	postByte := getExternalDataByteByUrl(url)
	post := &Post{}
	errUnM := json.Unmarshal(postByte, &post)
	if errUnM != nil {
		log.Println(errUnM)
	}
	return post
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
