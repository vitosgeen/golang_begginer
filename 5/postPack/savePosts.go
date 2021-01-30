package postPack

import (
	"log"
	"sync"
)

var CntPosts int = 0

func SaveToFilePostById(id int, mu *sync.Mutex) {
	mu.Lock()
	isSave := CntPosts < 5
	if isSave {
		post := GetPostById(id)
		SavePostToFile(CntPosts+1, *post)
		CntPosts++
	} else {
		log.Fatal("END")
	}
	mu.Unlock()
}