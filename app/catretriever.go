package app

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/asaintgenis/catapi/data/cat"
)

func searchImage() cat.Cat {
	randomInt := rand.Intn(2) + 1
	fileName := "./ressources/" + strconv.Itoa(randomInt) + ".jpg"
	catFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	content := make([]byte, 1024)
	catFile.Read(content)
	cat := cat.Cat{ID: "1", Pic: content}
	return cat
}
