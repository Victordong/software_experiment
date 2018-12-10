package model

type SupplyCommentModel struct {
	ID            uint         `json:"id"`
	CreatedAt     string       `json:"created_at"`
	User          *UserModel   `json:"user"`
	Username      string       `json:"username"`
	Content       string       `json:"content"`
	CommentedId   uint         `json:"commented_id"`
	CommentedName string       `json:"commented_name"`
	Supply        *SupplyModel `json:"supply"`
}
