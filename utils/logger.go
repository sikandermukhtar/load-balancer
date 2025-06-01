
// ### utils/logger.go
// Optional logger utility using the standard log package.

package utils

import "log"

func InitLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}