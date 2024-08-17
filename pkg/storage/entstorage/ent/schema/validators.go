package schema

import (
	"errors"
	"strings"
)

// Lowercase
func Lowercase(s string) error {
	if s != strings.ToLower(s) {
		return errors.New("only lowercase allowed")
	}
	return nil
}
