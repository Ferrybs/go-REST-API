package domain

type Blog struct {
	BlogID string `json:"blog_id"`
	UserID UserID `json:"user_id"`
}

type BlogDTO struct {
	UserID UserID `json:"user_id"`
}
