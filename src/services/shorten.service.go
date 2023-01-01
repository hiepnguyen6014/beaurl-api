package services

import (
	"fmt"

	"beaurl.vn/api/src/models"
	"github.com/google/uuid"
)

func ShortenLink(url string) (string, error) {

	uuid := uuid.NewString()[:8]

	err := models.Save(uuid, url)

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func GetLink(path string) string {

	shorten, err := models.FindByPath(path)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return shorten.Target
}
