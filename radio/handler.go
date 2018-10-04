package function

import (
	"fmt"
	"log"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	file, err := os.Open("./function/index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%q", data[:count])
	// return fmt.Sprintf("Welcome on RADIO, %s!", string(req))
}
