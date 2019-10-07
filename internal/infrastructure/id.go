package infrastructure

import (
	"crypto/md5"
	"fmt"
)

func GetRequestId(content []byte) string {
	return fmt.Sprintf("%x", md5.Sum(content))
}
