package api

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}

type CreateBlogRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
