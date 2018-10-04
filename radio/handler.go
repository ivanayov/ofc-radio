package function

import (
	"log"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	_, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
	// return fmt.Sprintf("Welcome on RADIO, %s!", string(req))
}
