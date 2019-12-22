package model

import (
	"GinHello/initDB"
	"log"
)

type Article struct {
	// id123
	ID int `json:"id"`
	// 类型
	Type string `json:"type"`
	// 内容
	Content string `json:"content"`
}

func (article Article) TableName() string {
	return "article"
}

func (article Article) Insert() int {
	create := initDB.Db.Create(&article)
	if create.Error != nil {
		log.Panicln("insert error is ", create.Error)
		return 0
	}
	return 1
}

func (article Article) FindById() Article {
	initDB.Db.First(&article, article.ID)
	return article
}

func (article Article) DeleteOne() {
	initDB.Db.Delete(article)
}

func (article Article) FindAll() []Article {
	var articles []Article
	initDB.Db.Find(&articles)
	return articles
}
