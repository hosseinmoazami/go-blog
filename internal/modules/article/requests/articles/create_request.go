package articles

import "mime/multipart"

type CreateRequest struct {
	Image   *multipart.FileHeader `form:"image" binding:"required"`
	Title   string                `form:"title" binding:"required,min=10,max=100"`
	Content string                `form:"content" binding:"required,min=10,max=600"`
}
