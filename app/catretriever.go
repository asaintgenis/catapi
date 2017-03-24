package app

import "github.com/asaintgenis/catapi/data/cat"
import "os"

func searchImage() cat.Cat {
	//randomInt := rand.Intn(2) + 1
	fileName := "./ressources/" + "2" + ".jpg"
	catFile, err := os.Open(fileName)
	if err != nil {
		panic("open fail")
	}
	content := make([]byte, 1024)
	catFile.Read(content)
	cat := cat.Cat{ID: "1", Pic: content}
	return cat
}
