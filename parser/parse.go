package parser

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/developer-guy/bookmarks-searcher/bookmark"
)

func Parse(jsonPath string) (interface{}, error) {
	var bookmarks bookmark.Bookmarks

	if jsonPath == "." {
		//setting current folder
		jsonPath = ""
	}

	file, err := os.Open(jsonPath + "Bookmarks.json")

	if err != nil {
		return nil, err
	}

	fmt.Println("File opened successuflly", file.Name())

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &bookmarks)

	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}
