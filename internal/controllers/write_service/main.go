package controllers

import (
	"math/rand"

	"github.com/w1lay/Simple-URL-Shortener/internal/dtos"
	"github.com/w1lay/Simple-URL-Shortener/internal/repositories"
)

const (
	Charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	LinkLength = 8
)

type WriteServiceController struct {
	repo *repositories.LinkPostgresRepository
}

func NewWriteServiceController(repo *repositories.LinkPostgresRepository) *WriteServiceController {
	return &WriteServiceController{repo: repo}
}

func generateLongLink() string {
	bytes := make([]byte, LinkLength, LinkLength)
	for index := range bytes {
		bytes[index] = Charset[rand.Intn(len(Charset))]
	}

	return string(bytes)
}

func (c *WriteServiceController) CreateLink(longLink string) (*dtos.Link, error) {
	dto := dtos.Link{
		ShortLink: generateLongLink(),
		LongLink:  longLink,
	}

	err := c.repo.SaveLinkRecord(dto)
	return &dto, err
}
