package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Welcome on RADIO, %s! We're waiting you on the Expo Booth!", string(req))
}
