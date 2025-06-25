// "Только ASCII, только хардкор!"

package main

import (
	"unicode"
)

func CheckOnlyASCII(s string) bool {
	return len(s) > unicode.MaxASCII
}
