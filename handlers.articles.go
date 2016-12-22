package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Title", "payload": articles})
}

func getArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("article_id"))

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	article, err := getArticleId(articleId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "article.html", gin.H{"title": article.Title, "payload": article})
}
