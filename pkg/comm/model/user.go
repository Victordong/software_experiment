package model

type UserModel struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordHash string `json:"-"`
	Name         string `json:"name"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	IconUrl      string `json:"icon_url"`
	Role         string `json:"role"`
	Address      string `json:"address"`
	QqNumber     string `json:"qq_number"`
	Information  string `json:"information"`
}

type ChangePasswordUsername struct {
	Username string `json:"username"`
}

type ChangePasswordModel struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	Session     string `json:"session"`
}

type CaptchaModel struct {
	CaptchaId string `json:"captcha_id"`
	Result    string `json:"result,omitempty" binding:"required"`
}
