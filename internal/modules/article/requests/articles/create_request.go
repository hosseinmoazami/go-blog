package articles

type CreateRequest struct {
	Title   string `form:"title" binding:"required,min=10,max=100"`
	Content string `form:"content" binding:"required,min=10,max=600"`
}
