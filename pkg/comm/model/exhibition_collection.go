package model

type ExhibitionCollectionModel struct {
	ID            uint             `json:"id"`
	CreatedAt     string           `json:"created_at"`
	User          *UserModel       `json:"user"`
	CollectedName string           `json:"collected_name"`
	Username      string           `json:"username"`
	CollectedId   uint             `json:"collected_id"`
	Exhibition    *ExhibitionModel `json:"exhibition"`
}
