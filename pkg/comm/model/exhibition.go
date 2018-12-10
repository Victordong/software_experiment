package model

type ExhibitionModel struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Theme     string `json:"theme"`
	Name      string `json:"name"`
	VisitNum  int    `json:"visit_num"`
}

type ExhibitionListNodeModel struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Theme     string `json:"theme"`
	VisitNum  int    `json:"visit_num"`
}
