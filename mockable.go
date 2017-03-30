package mockable

import (
	"os"
)

func Mocked() bool {
	if len(os.Getenv("MOCKABLE")) > 0 {
		return true
	}

	return false
}
