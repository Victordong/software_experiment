package model

type ExhibitionModel struct {
	ID        uint       `json:"id"`
	CreatedAt string     `json:"created_at"`
	User      *UserModel `json:"user"`
	Username  string     `json:"username"`
	Content   string     `json:"content,omitempty"`
	Theme     string     `json:"theme"`
	Name      string     `json:"name"`
	VisitNum  int        `json:"visit_num"`
}
