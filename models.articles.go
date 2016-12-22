package main

import (
	"errors"
)

type article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{1, "title1", "content"},
	{2, "title2", "content2"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleId(articleId int) (*article, error) {
	for _, articleStruct := range articleList {
		if articleId == articleStruct.Id {
			return &articleStruct, nil
		}
	}

	return nil, errors.New("article not found")
}
