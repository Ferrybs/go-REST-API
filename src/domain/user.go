package domain

type UserID string

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginView struct {
	AccessToken string   `json:"access_token"`
	User        UserView `json:"user"`
}

type UserView struct {
	UserID   UserID `json:"user_id"`
	Username string `json:"username"`
}

type User struct {
	UserID   UserID `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
