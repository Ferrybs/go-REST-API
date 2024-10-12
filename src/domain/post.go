package domain

type PostID string

type PostDTO struct {
	Title   string `json:"title"`
	UserID  UserID `json:"user_id"`
	Content string `json:"content"`
}

type PostView struct {
	ID        string `json:"id"`
	UserID    UserID `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Post struct {
	ID        string `json:"id"`
	UserID    UserID `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
