package api

import (
	"blog/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "blog/api/docs" // for swagger
)

type handler struct {
	storage *storage.DBManager
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      localhost:8000
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
