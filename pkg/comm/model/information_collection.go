package model

type InformationCollectionModel struct {
	ID            uint              `json:"id"`
	CreatedAt     string            `json:"created_at"`
	User          *UserModel        `json:"user"`
	CollectedName string            `json:"collected_name"`
	Username      string            `json:"username"`
	CollectedId   uint              `json:"collected_id"`
	Information   *InformationModel `json:"information"`
}
