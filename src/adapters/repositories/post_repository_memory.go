package repositories

import (
	"blog/api/src/domain"
	"time"

	"github.com/google/uuid"
)

type PostRepositoryMemory struct {
	Post []domain.Post
}

func NewPostRepositoryMemory() IPostRepository {
	return &PostRepositoryMemory{}
}

func (b *PostRepositoryMemory) CreatePost(PostDTO *domain.PostDTO) (*domain.Post, error) {
	Post := domain.Post{
		ID:        uuid.New().String(),
		UserID:    PostDTO.UserID,
		Title:     PostDTO.Title,
		Content:   PostDTO.Content,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	b.Post = append(b.Post, Post)
	return &Post, nil
}

func (b *PostRepositoryMemory) FindAllPostByUserID(userID domain.UserID, page int, limit int) ([]domain.Post, error) {
	Post := []domain.Post{}
	for _, P := range b.Post {
		if P.UserID == userID {
			Post = append(Post, P)
		}
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= len(Post) {
		return []domain.Post{}, nil
	}

	if endIndex > len(Post) {
		endIndex = len(Post)
	}

	return Post[startIndex:endIndex], nil
}

func (b *PostRepositoryMemory) FindPostByID(postID domain.PostID) (*domain.Post, error) {
	for _, Post := range b.Post {
		if Post.ID == string(postID) {
			return &Post, nil
		}
	}
	return nil, nil
}

func (b *PostRepositoryMemory) FindAllPost(page int, limit int) ([]domain.Post, error) {
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= len(b.Post) {
		return []domain.Post{}, nil
	}

	if endIndex > len(b.Post) {
		endIndex = len(b.Post)
	}
	return b.Post[:endIndex], nil
}

func (b *PostRepositoryMemory) DeletePostByPostID(postID domain.PostID) error {
	for i, P := range b.Post {
		if P.ID == string(postID) {
			b.Post = append(b.Post[:i], b.Post[i+1:]...)
			return nil
		}
	}
	return nil
}
