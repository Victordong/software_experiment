package model

type ExhibitionCommentModel struct {
	ID            uint             `json:"id"`
	CreatedAt     string           `json:"created_at"`
	User          *UserModel       `json:"user"`
	Username      string           `json:"username"`
	Content       string           `json:"content"`
	CommentedId   uint             `json:"commented_id"`
	CommentedName string           `json:"commented_name"`
	Exhibition    *ExhibitionModel `json:"exhibition"`
}
