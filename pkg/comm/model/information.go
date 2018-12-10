package model

type InformationModel struct {
	ID        uint       `json:"id"`
	CreatedAt string     `json:"created_at"`
	Username  string     `json:"username"`
	User      *UserModel `json:"user"`
	Content   string     `json:"content,omitempty"`
	Name      string     `json:"name"`
	VisitNum  int        `json:"visit_num"`
}
