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

	r.GET("/blogs/:id", h.GetBlog)
	r.GET("/blogs", h.GetBlogs)
	r.POST("/blogs", h.CreateBlog)
	r.PUT("/blogs/:id", h.UpdateBlog)
	r.DELETE("/blogs/:id", h.DeleteBlog)

	return r
}
