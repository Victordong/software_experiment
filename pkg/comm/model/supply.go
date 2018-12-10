package model

import (
	"time"
)

type SupplyModel struct {
	ID         uint       `json:"id"`
	CreatedAt  string     `json:"created_at"`
	User       *UserModel `json:"user"`
	Username   string     `json:"username"`
	Name       string     `json:"name"`
	Content    string     `json:"content,omitempty"`
	VisitNum   int        `json:"visit_num"`
	Type       string     `json:"type"`
	ExpiryDate time.Time  `json:"expiry_date"`
}
