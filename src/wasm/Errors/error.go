package errors

import (
	"fmt"
	"os"
)

func AssertError(boo bool, text string) {
	if boo {
		fmt.Println(text)
		os.Exit(1)
	}
}
