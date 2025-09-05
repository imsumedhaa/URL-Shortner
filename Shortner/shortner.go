package shortner

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres"
)

type URL struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	postgres.Client
}


func Generator(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))

	data := hasher.Sum(nil)
	abc := hex.EncodeToString(data)

	return abc[:8]
}
