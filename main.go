package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type URL struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

var urlDb = make(map[string]URL)

func generator(OriginalURL string)string{
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println(hasher)

	data:=hasher.Sum(nil)
	//fmt.Println(data)
	abc := hex.EncodeToString(data)
	//fmt.Println(abc)

	return abc[:8]
}

func Shortner(originalUrl string)string{
	shortUrl := generator(originalUrl)
	id := shortUrl
	urlDb[id] = URL{
		ID: id,
		OriginalURL: originalUrl,
		ShortURL: shortUrl,
	}
	fmt.Println(urlDb)
	return ""
}

func main() {
	fmt.Println("URL Shortner")
	OriginalURL := "https://kubernetes.io/docs/concepts/services-networking/ingress/"
	fmt.Println(Shortner(OriginalURL))	
}
