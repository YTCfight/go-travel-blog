package model

import (
	"github.com/gin-gonic/gin"
	"go-travel-blog/pkg/app"
	"go-travel-blog/pkg/errcode"
)

type Article struct {
	*Model
	Title           string `json:"title"`
	Desc            string `json:"desc"`
	Content         string `json:"content"`
	ConvertImageUrl string `json:"convert_image_url"`
	State           uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
