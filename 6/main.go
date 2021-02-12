package main

import (
	"fmt"
	"golang_begginer/6/postPack"
)

func main() {
	posts := postPack.GetPostsByUserId(7)
	for _, post := range *posts {
		go savePostUser(post)
	}
	go postPack.ChanelReadPostWorker()
	go postPack.ChanelReadCommentWorker()

	//I don't understand it. How do without it? I have not found answer yet)))
	var input string
	fmt.Scanln(&input)
}
func savePostUser(post postPack.Post) {
	postPack.ChanelWritePostWorker(post)
	go saveCommentsByPostId(post.ID)
}
func saveCommentsByPostId(id int) {
	comments := postPack.GetCommentsByPostId(id)
	for _, comment := range *comments {
		postPack.ChanelWriteCommentWorker(comment)
	}
}
