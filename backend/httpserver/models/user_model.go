package models

type UserModel struct {
	BaseModel
	Full_name string `json:"full_name"`
	Email     string `gorm:"uniqueIndex" json:"email"`
	Password  string `json:"password"`
}

func (UserModel) TableName() string {
	return "public.User"
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
