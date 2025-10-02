package structs

// Struct for comment responses
type CommentResponse struct {
	Id        uint   `json:"id"`
	Content   string `json:"content"`
	PostID    uint   `json:"post_id"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Request for creating a comment
type CommentCreateRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// Request for updating a comment
type CommentUpdateRequest struct {
	Content string `json:"content" binding:"required"`
}
