package util

import (
	"log"

	"github.com/joho/godotenv"
)

var validMimeTypes = []string{"image/jpg", "image/jpeg", "image/png", "application/pdf"}

func LoadDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}
}

func IsValidMimeType(mime string) bool {
	for _, validMimeType := range validMimeTypes {
		if mime == validMimeType {
			return true
		}
	}
	return false
}

func GetFileExtension(mime string) string {
	switch mime {
	case "image/jpg", "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	case "application/pdf":
		return "pdf"
	default:
		return ""
	}
}
