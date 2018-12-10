package model

type SupplyCollectionModel struct {
	ID            uint   `json:"id"`
	CreatedAt     string `json:"created_at"`
	CollectedName string `json:"collected_name"`
	Username      string `json:"username"`
	CollectedId   uint   `json:"collected_id"`
}
