package model

type SupplyModel struct {
	ID         uint   `json:"id"`
	CreatedAt  string `json:"created_at"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	VisitNum   int    `json:"visit_num"`
	Type       int    `json:"type"`
	ExpiryDate string `json:"expiry_date"`
}

type SupplyListNodeModel struct {
	ID         uint   `json:"id"`
	CreatedAt  string `json:"created_at"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	VisitNum   int    `json:"visit_num"`
	Type       string `json:"type"`
	ExpiryDate string `json:"expiry_date"`
}
