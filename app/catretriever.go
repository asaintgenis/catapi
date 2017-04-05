package app

import (
	"encoding/base64"
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
		return cat.Cat{ID: "", Pic: "", Err: "can't open file"}, err
	}

	defer catFile.Close()

	fInfo, _ := catFile.Stat()
	var size int64 = fInfo.Size()
	content := make([]byte, size)
	catFile.Read(content)
	imgBase64Str := base64.StdEncoding.EncodeToString(content)
	cat := cat.Cat{ID: "1", Pic: imgBase64Str}
	return cat, nil
}
