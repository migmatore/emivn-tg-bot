package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateLink() string {
	hash := sha256.Sum256([]byte(time.Now().Format("2006-01-02")))

	return hex.EncodeToString(hash[:5])
}
