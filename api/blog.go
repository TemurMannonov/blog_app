package api

import (
	"blog/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBlog(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse param",
		})
		return
	}

	blog, err := h.storage.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get blog",
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (h *handler) CreateBlog(ctx *gin.Context) {
	var b storage.Blog

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed bind request body",
		})
		return
	}

	blog, err := h.storage.Create(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create blog",
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}
