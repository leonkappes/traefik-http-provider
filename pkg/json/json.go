package json

import (
	"io/ioutil"
	"log"
	"os"
)

var ValueChanged = true
var cache = map[string][]byte{}

func ReadJsonFile(file string) []byte {
	if !ValueChanged {
		return cache[file]
	}

	jsonFile, err := os.Open(file)
	if err != nil {
		log.Printf("Error reading file %s: %s", file, err)
	}

	defer jsonFile.Close()

	jsonData, _ := ioutil.ReadAll(jsonFile)

	cache[file] = jsonData

	return jsonData
}
