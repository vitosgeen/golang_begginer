package postPack

import (
	"log"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var ChanelComment = make(chan Comment)

func loadCommentById(id int) Comment {
	var commentData Comment
	db := dbConnInit()
	selDB, err := db.Query("SELECT  `id`, `post_id`, `name`, `email`, `body`  FROM comment WHERE id = ? LIMIT 1", id)
	if err != nil {
		log.Println("loadCommentById SELECT ERROR: ", err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&commentData.ID, &commentData.PostID, &commentData.Name, &commentData.Email, &commentData.Body)
		if err != nil {
			log.Println("loadCommentById row  ERROR: ", err.Error())
		}
	}
	defer selDB.Close()
	defer db.Close()

	return commentData
}

func loadCommentByComment(comment Comment) Comment {
	var commentData Comment
	db := dbConnInit()
	selDB, err := db.Query("SELECT `id`, `post_id`, `name`, `email`, `body` FROM comment WHERE `post_id` = ? AND `name` = ? AND `email` = ? AND `body` = ? LIMIT 1", comment.PostID, comment.Name, comment.Email, comment.Body)
	if err != nil {
		log.Println("loadCommentByComment SELECT ERROR: ", err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&commentData.ID, &commentData.PostID, &commentData.Name, &commentData.Email, &commentData.Body)
		if err != nil {
			log.Println("loadCommentByComment row  ERROR: ", err.Error())
		}
	}
	selDB.Close()
	defer db.Close()

	return commentData
}

func SaveComment(comment Comment) Comment {
	loadComment := loadCommentById(comment.ID)
	var saveComment Comment
	if loadComment.ID != 0 {
		saveComment = updateComment(comment)
	} else {
		saveComment = insertComment(comment)
	}

	return saveComment
}

func insertComment(comment Comment) Comment {
	db := dbConnInit()
	insertSQL, err := db.Prepare("INSERT INTO `comment` (`id`, `post_id`, `name`, `email`, `body`) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("insertComment INSERT Prepare ERROR: ", err.Error())
		return comment
	}
	_, err = insertSQL.Exec(comment.ID, comment.PostID, comment.Name, comment.Email, comment.Body)
	if err != nil {
		log.Println("insertComment INSERT ERROR: ", err.Error())
		return comment
	}
	defer insertSQL.Close()
	defer db.Close()

	return loadCommentByComment(comment)
}

func updateComment(comment Comment) Comment {
	db := dbConnInit()
	updateSQL, err := db.Prepare("UPDATE `comment` SET `post_id` = ?, `name` = ?, `email` = ?, `body` = ? WHERE `id` = ?")
	if err != nil {
		log.Println("updateComment UPDATE Prepare ERROR: ", err.Error())
		return comment
	}
	_, err = updateSQL.Exec(comment.PostID, comment.Name, comment.Email, comment.Body, comment.ID)
	if err != nil {
		log.Println("updateComment UPDATE ERROR: ", err.Error())
		return comment
	}
	defer updateSQL.Close()
	defer db.Close()

	return comment
}

func ChanelReadCommentWorker() {
	for {
		comment, ok := <-ChanelComment
		if ok == false {
			break
		} else {
			SaveComment(comment)
		}
	}
}
func ChanelWriteCommentWorker(comment Comment) {
	ChanelComment <- comment
}
