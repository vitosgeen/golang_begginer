package postPack

import (
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

func loadUserById(id int) User {
	var userData User
	db := dbConnInit()
	selDB, err := db.Query("SELECT `id`, `name`, `username`, `email`, `phone`, `website` FROM user WHERE id = ? LIMIT 1", id)
	if err != nil {
		log.Println("loadUserById SELECT ERROR: ", err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&userData.ID, &userData.Name, &userData.Username, &userData.Email, &userData.Phone, &userData.Website)
		if err != nil {
			log.Println("loadUserById row  ERROR: ", err.Error())
		}
	}
	defer selDB.Close()
	defer db.Close()

	return userData
}

func loadUserByUser(user User) User {
	var userData User
	db := dbConnInit()
	selDB, err := db.Query("SELECT `id`, `name`, `username`, `email`, `phone`, `website` FROM `user` WHERE `name` = ? AND `username` = ? AND `email` = ? AND `phone` = ? AND `website` = ? LIMIT 1", user.Name, user.Username, user.Email, user.Phone, user.Website)
	if err != nil {
		log.Println("loadUserByUser SELECT ERROR: ", err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&userData.ID, &userData.Name, &userData.Username, &userData.Email, &userData.Phone, &userData.Website)
		if err != nil {
			log.Println("loadUserByUser row  ERROR: ", err.Error())
		}
	}
	selDB.Close()
	defer db.Close()

	return userData
}

func saveUser(user User) User {
	loadUser := loadUserById(user.ID)
	var saveUser User
	if loadUser.ID != 0 {
		saveUser = updateUser(user)
	} else {
		saveUser = insertUser(user)
	}

	return saveUser
}

func insertUser(user User) User {
	db := dbConnInit()
	insertSQL, err := db.Prepare("INSERT INTO `user` (`name`, `username`, `email`, `phone`, `website`) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("insertUser INSERT Prepare ERROR: ", err.Error())
		return user
	}
	_, err = insertSQL.Exec(user.ID, user.Name, user.Username, user.Email, user.Phone, user.Website)
	if err != nil {
		log.Println("insertUser INSERT ERROR: ", err.Error())
		return user
	}
	defer insertSQL.Close()
	defer db.Close()

	return loadUserByUser(user)
}

func updateUser(user User) User {
	db := dbConnInit()
	updateSQL, err := db.Prepare("UPDATE `user` SET `name` = ?, `username` = ?, `email` = ?, `phone` = ?, `website` = ? WHERE `id` = ?")
	if err != nil {
		log.Println("updateUser UPDATE Prepare ERROR: ", err.Error())
		return user
	}
	_, err = updateSQL.Exec(user.Name, user.Username, user.Email, user.Phone, user.Website, user.ID)
	if err != nil {
		log.Println("updateUser UPDATE ERROR: ", err.Error())
		return user
	}
	defer updateSQL.Close()
	defer db.Close()

	return user
}
