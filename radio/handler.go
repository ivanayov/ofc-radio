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

	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	filesize := fileinfo.Size()

	data := make([]byte, filesize)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
	// return fmt.Sprintf("Welcome on RADIO, %s!", string(req))
}
