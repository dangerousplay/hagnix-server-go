package service

import (
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/modelxml"
)

var news = &NewsService{}

type NewsService struct{}

func (news *NewsService) GetNews() ([]modelxml.NewsItemXML, error) {
	var newsM []models.News
	var newsXML = []modelxml.NewsItemXML{}

	err := database.GetDBEngine().Limit(10).Find(&newsM)

	if err != nil {
		return nil, err
	}

	for _, v := range newsM {
		newsXML = append(newsXML, modelxml.NewsItemXML{
			Icon:    v.Icon,
			Link:    v.Icon,
			TagLine: v.Text,
			Date:    v.Date.Unix(),
			Title:   v.Title,
		})
	}

	return newsXML, err
}

func GetNewsService() *NewsService {
	return news
}
