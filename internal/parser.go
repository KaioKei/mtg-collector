package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func readFile(path string) []byte {

	// open file
	f, err := os.Open(path)
	defer f.Close()
	FatalIfErr(err, fmt.Sprint("Cannot open file:", path))

	// reading large files with a loop
	log.Println("Read file:", path)
	finalBuffer := make([]byte, 0)
	br := bufio.NewReader(f)
	for {
		// reading 1024 bytes (100Ko) at each iteration
		largeBuffer := make([]byte, 102400)
		bytesRead, err := br.Read(largeBuffer)
		// checking if bytes were read
		if bytesRead == 0 {
			if err == io.EOF {
				// stop if EOF is reached
				break
			}
			FatalIfErr(err, fmt.Sprint("Could not read all bytes for file:", path))
		}
		// concatenate largeBuffer to the bytes read only
		largeBuffer = largeBuffer[:bytesRead]
		// append to the final buffer
		finalBuffer = append(finalBuffer, largeBuffer...)
	}

	log.Printf("Reading file ended with %d bytes: %s", len(finalBuffer), path)
	return finalBuffer
}

func getKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

func ParseFile(path string) {
	contentBytes := readFile(path)
	var jsonData map[string]interface{}
	if err := json.Unmarshal(contentBytes, &jsonData); err != nil {
		log.Fatal(err)
	}
	// show all
	//log.Println(jsonData)
	// show keys

	log.Println(getKeys(jsonData))
}
