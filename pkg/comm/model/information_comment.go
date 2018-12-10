package model

type InformationCommentModel struct {
	ID            uint   `json:"id"`
	CreatedAt     string `json:"created_at"`
	Username      string `json:"username"`
	Content       string `json:"content"`
	CommentedId   uint   `json:"commented_id"`
	CommentedName string `json:"commented_name"`
}
