package app

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/asaintgenis/catapi/data/cat"
)

func createSeachURL() string {

	searchURL := []string{"https://www.googleapis.com/customsearch/v1?"}
	query := make(map[string]string)
	query["q"] = "lolcat"
	query["searchType"] = "image"
	query["fields"] = "items(link)"
	// add your google api keys
	query["cx"] = ""
	query["key"] = ""
	// override defaults
	if os.Getenv("GOOGLE_CX") != "" {
		query["cx"] = os.Getenv("GOOGLE_CX")
	}
	if os.Getenv("GOOGLE_KEY") != "" {
		query["key"] = os.Getenv("GOOGLE_KEY")
	}

	for k, v := range query {
		searchURL = append(searchURL, k, "=", v, "&")
	}

	return strings.Join(searchURL, "")

}

func searchImage() cat.Cat {
	imageSearch := createSeachURL()

	request, err := http.Get(imageSearch)
	if err != nil {
		panic(err)
	}
	if request.StatusCode != http.StatusOK {
		panic("http status code not 200")
	}
	defer request.Body.Close()
	content, _ := ioutil.ReadAll(request.Body)

	cat := cat.Cat{ID: "1", Pic: content}

	return cat
}
