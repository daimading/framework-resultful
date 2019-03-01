package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Sha1Hash(content string) string {
	hash := sha1.New()
	io.WriteString(hash, content)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
