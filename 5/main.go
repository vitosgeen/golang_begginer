package main

import (
	"fmt"
	"golang_begginer/5/postPack"
	"sync"
)

func main() {
	fmt.Println("START")
	postsID := postPack.GetPostsId()
	mu := new(sync.Mutex)
	for _, post := range *postsID {
		go postPack.SaveToFilePostById(post.Id, mu)
	}
	var input string
	fmt.Scanln(&input)
}