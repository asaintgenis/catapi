package app

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/asaintgenis/catapi/data/cat"
)

func searchImage(catID string) (cat.Cat, error) {
	if len(catID) == 0 {
		catID = strconv.Itoa(rand.Intn(2) + 1)
	}

	fileName := "./ressources/" + catID + ".jpg"
	catFile, err := os.Open(fileName)
	if err != nil {
		return cat.Cat{ID: "", Pic: nil, Err: "can't open file"}, err
	}
	content := make([]byte, 1024)
	catFile.Read(content)
	cat := cat.Cat{ID: "1", Pic: content}
	return cat, nil
}
