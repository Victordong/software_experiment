package model

type InformationCommentModel struct {
	ID            uint              `json:"id"`
	CreatedAt     string            `json:"created_at"`
	User          *UserModel        `json:"user"`
	Username      string            `json:"username"`
	Information   *InformationModel `json:"information"`
	Content       string            `json:"content"`
	CommentedId   uint              `json:"commented_id"`
	CommentedName string            `json:"commented_name"`
}
