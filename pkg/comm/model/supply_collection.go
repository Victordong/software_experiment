package model

type SupplyCollectionModel struct {
	ID            uint         `json:"id"`
	CreatedAt     string       `json:"created_at"`
	User          *UserModel   `json:"user"`
	CollectedName string       `json:"collected_name"`
	Username      string       `json:"username"`
	CollectedId   uint         `json:"collected_id"`
	Supply        *SupplyModel `json:"supply"`
}
