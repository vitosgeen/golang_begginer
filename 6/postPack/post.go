package postPack

import (
	"log"
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

var ChanelPost = make(chan Post)

func loadPostById(id int) Post {
	var postData Post
	db := dbConnInit()
	selDB, err := db.Query("SELECT `id`, `title`, `user_id`, `body` FROM post WHERE id = ? LIMIT 1", id)
	if err != nil {
		log.Println("loadPostById SELECT ERROR: ", err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&postData.ID, &postData.Title, &postData.UserID, &postData.Body)
		if err != nil {
			log.Println("loadPostById row  ERROR: ", err.Error())
		}
	}
	defer selDB.Close()
	defer db.Close()

	return postData
}

func loadPostByPost(post Post) Post {
	var postData Post
	db := dbConnInit()
	selDB, err := db.Query("SELECT `id`, `title`, `user_id`, `body` FROM post WHERE `title` = ? AND `user_id` = ? AND `body` = ? LIMIT 1", post.Title, post.UserID, post.Body)
	if err != nil {
		log.Println("loadPostByPost SELECT ERROR: ", err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&postData.ID, &postData.Title, &postData.UserID, &postData.Body)
		if err != nil {
			log.Println("getAppEntity row  ERROR: ", err.Error())
		}
	}
	selDB.Close()
	defer db.Close()

	return postData
}

func SavePost(post Post) Post {
	loadPost := loadPostById(post.ID)
	log.Println(loadPost)
	var savePost Post
	if loadPost.ID != 0 {
		savePost = updatePost(post)
	} else {
		savePost = insertPost(post)
	}

	return savePost
}

func insertPost(post Post) Post {
	db := dbConnInit()
	insertSQL, err := db.Prepare("INSERT INTO `post`(`id`, `title`, `user_id`, `body`) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Println("insertPost INSERT Prepare ERROR: ", err.Error())
		return post
	}
	_, err = insertSQL.Exec(post.ID, post.Title, post.UserID, post.Body)
	if err != nil {
		log.Println("insertPost INSERT ERROR: ", err.Error())
		return post
	}
	defer insertSQL.Close()
	defer db.Close()

	return loadPostByPost(post)
}

func updatePost(post Post) Post {
	db := dbConnInit()
	updateSQL, err := db.Prepare("UPDATE `post` SET `title` = ?, `user_id` = ?, `body` = ? WHERE `id` = ?")
	if err != nil {
		log.Println("updatePost UPDATE Prepare ERROR: ", err.Error())
		return post
	}
	_, err = updateSQL.Exec(post.Title, post.UserID, post.Body, post.ID)
	if err != nil {
		log.Println("updatePost UPDATE ERROR: ", err.Error())
		return post
	}
	defer updateSQL.Close()
	defer db.Close()

	return post
}

func ChanelReadPostWorker() {
	for {
		post, ok := <-ChanelPost
		if ok == false {
			break
		} else {
			SavePost(post)
		}
	}
}
func ChanelWritePostWorker(post Post) {
	ChanelPost <- post
}
