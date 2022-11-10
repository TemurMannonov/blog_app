package api

import (
	"blog/storage"

	"github.com/gin-gonic/gin"
)

type handler struct {
	storage *storage.DBManager
}

func NewServer(storage *storage.DBManager) *gin.Engine {
	r := gin.Default()

	h := handler{
		storage: storage,
	}

	r.GET("/blog/:id", h.GetBlog)
	r.POST("/blog", h.CreateBlog)

	return r
}
