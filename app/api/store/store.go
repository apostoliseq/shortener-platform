package store

import (
	"math/rand"
	"time"
)

// URLRecord represents a stored URL mapping
type URLRecord struct {
	ID        int
	Code      string
	LongURL   string
	CreatedAt time.Time
}

// in-memory store: code → URLRecord
var records = map[string]URLRecord{}
var nextID = 1

func Save(longURL string) URLRecord {
	code := generateCode(6)
	record := URLRecord{
		ID:        nextID,
		Code:      code,
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}
	records[code] = record
	nextID++
	return record
}

func GetByCode(code string) (URLRecord, bool) {
	record, ok := records[code]
	return record, ok
}

func generateCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}
