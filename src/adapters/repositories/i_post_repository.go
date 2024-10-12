package repositories

import "blog/api/src/domain"

type IPostRepository interface {
	CreatePost(Post *domain.PostDTO) (*domain.Post, error)
	FindAllPostByUserID(userID domain.UserID, page int, limit int) ([]domain.Post, error)
	FindPostByID(postID domain.PostID) (*domain.Post, error)
	FindAllPost(page int, limit int) ([]domain.Post, error)
	DeletePostByPostID(postID domain.PostID) error
}
