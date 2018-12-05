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
	var newsXML []modelxml.NewsItemXML

	err := database.GetDBEngine().Limit(10).Find(&newsM)

	if err != nil {
		return nil, err
	}

	if len(newsM) < 1 {
		newsM = getDefaultNews()
	}

	for _, v := range newsM {
		newsXML = append(newsXML, modelxml.NewsItemXML{
			Icon:    v.Icon,
			Link:    v.Link,
			TagLine: v.Text,
			Date:    v.Date.Unix(),
			Title:   v.Title,
		})
	}

	if len(newsM) < 1 {

	}

	return newsXML, err
}

func getDefaultNews() []models.News {
	return []models.News{{Title: "Servidor Hagnix", Text: "Lançamento do ano"}}
}

func GetNewsService() *NewsService {
	return news
}
