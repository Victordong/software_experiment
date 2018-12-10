package model

type InformationModel struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Name      string `json:"name"`
	VisitNum  int    `json:"visit_num"`
}

type InformationListNodeModel struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	VisitNum  int    `json:"visit_num"`
}
