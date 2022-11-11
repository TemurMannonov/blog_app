package api

import (
	"blog/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router /blogs/{id} [get]
// @Summary Get blog by id
// @Description Get blog by id
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} storage.Blog
// @Failure 500 {object} ResponseError
func (h *handler) GetBlog(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	blog, err := h.storage.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// @Router /blogs [get]
// @Summary Get blogs
// @Description Get blogs
// @Tags blog
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param author query string false "Author"
// @Param title query string false "Title"
// @Success 200 {object} storage.GetBlogsResult
// @Failure 500 {object} ResponseError
func (h *handler) GetBlogs(ctx *gin.Context) {
	queryParams, err := validateGetBlogsQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.storage.GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetBlogsQuery(ctx *gin.Context) (*storage.GetBlogsQueryParam, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &storage.GetBlogsQueryParam{
		Limit:  int32(limit),
		Page:   int32(page),
		Author: ctx.Query("author"),
		Title:  ctx.Query("title"),
	}, nil
}

// @Router /blogs [post]
// @Summary Create a blog
// @Description Create a blog
// @Tags blog
// @Accept json
// @Produce json
// @Param blog body CreateBlogRequest true "Blog"
// @Success 200 {object} storage.Blog
// @Failure 500 {object} ResponseError
func (h *handler) CreateBlog(ctx *gin.Context) {
	var b storage.Blog

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	created, err := h.storage.Create(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, created)
}

// @Router /blogs/{id} [put]
// @Summary Update a blog
// @Description Update a blog
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param blog body CreateBlogRequest true "Blog"
// @Success 200 {object} storage.Blog
// @Failure 500 {object} ResponseError
func (h *handler) UpdateBlog(ctx *gin.Context) {
	var b storage.Blog

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	b.ID = id

	updated, err := h.storage.Update(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

// @Router /blogs/{id} [delete]
// @Summary Delete a blog
// @Description Delete a blog
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} ResponseOK
// @Failure 500 {object} ResponseError
func (h *handler) DeleteBlog(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.storage.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted!",
	})
}
