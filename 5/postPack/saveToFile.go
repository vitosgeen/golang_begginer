package postPack

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const pathStoreBaseDir = "./5/storage"
const pathStoreBaseDirPosts = "./5/storage/posts"

func SavePostToFile(num int, post Post) {
	prepareSavePostToFile()
	postByte, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}
	filename := pathStoreBaseDirPosts + "/" + strconv.Itoa(num) + ".txt"
	//SavePostToFileIoutil(filename, postByte)
	SavePostToFileBufio(filename, postByte)
}

func SavePostToFileIoutil(filename string, postByte []byte) {
	err := ioutil.WriteFile(filename, postByte, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func SavePostToFileBufio(filename string, postByte []byte)  {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	err = file.Chmod(os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
    _, err = writer.Write(postByte)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

func prepareSavePostToFile()  {
	_, err := os.Stat(pathStoreBaseDir)
	if err != nil {
		os.MkdirAll(pathStoreBaseDir, os.ModePerm)
	}
	_, err = os.Stat(pathStoreBaseDirPosts)
	if err != nil {
		os.MkdirAll(pathStoreBaseDirPosts, os.ModePerm)
	}
}