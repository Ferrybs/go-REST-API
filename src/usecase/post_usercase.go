package usecase

import (
	"blog/api/src/adapters/repositories"
	"blog/api/src/domain"
	"errors"
)

type PostUsecase struct {
	PostRepository repositories.IPostRepository
	userRepository repositories.IUserRepository
}

func NewPostUsecase(PostRepository repositories.IPostRepository, userRepository repositories.IUserRepository) *PostUsecase {
	return &PostUsecase{PostRepository: PostRepository, userRepository: userRepository}
}

func (b *PostUsecase) CreatePost(PostDTO *domain.PostDTO) (*domain.PostView, error) {
	Post, err := b.PostRepository.CreatePost(PostDTO)
	if err != nil {
		return nil, err
	}
	PostView := domain.PostView{
		ID:        Post.ID,
		UserID:    Post.UserID,
		Title:     Post.Title,
		Content:   Post.Content,
		CreatedAt: Post.CreatedAt,
		UpdatedAt: Post.UpdatedAt,
	}

	return &PostView, nil
}

func (b *PostUsecase) FindAllPostByUsername(username string, page int, limit int) ([]domain.PostView, error) {
	user, err := b.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	Post, err := b.PostRepository.FindAllPostByUserID(user.UserID, page, limit)
	if err != nil {
		return nil, err
	}
	PostViews := make([]domain.PostView, len(Post))
	for i, Post := range Post {
		PostViews[i] = domain.PostView(Post)
	}
	return PostViews, nil
}

func (b *PostUsecase) FindPostByID(postID domain.PostID) (*domain.PostView, error) {
	Post, err := b.PostRepository.FindPostByID(postID)
	if err != nil {
		return nil, err
	}
	if Post == nil {
		return nil, nil
	}
	PostView := domain.PostView(*Post)
	return &PostView, nil
}

func (b *PostUsecase) FindAllPost(page int, limit int) ([]domain.PostView, error) {
	Post, err := b.PostRepository.FindAllPost(page, limit)
	if err != nil {
		return nil, err
	}
	PostViews := make([]domain.PostView, len(Post))
	for i, Post := range Post {
		PostViews[i] = domain.PostView(Post)
	}
	return PostViews, nil
}

func (b *PostUsecase) DeletePostByPostID(userID domain.UserID, postID domain.PostID) error {
	Post, err := b.PostRepository.FindPostByID(postID)
	if err != nil {
		return err
	}
	if Post == nil {
		return errors.New("post not found")
	}
	if Post.UserID != userID {
		return errors.New("you are not the owner of this post")
	}
	err = b.PostRepository.DeletePostByPostID(postID)
	if err != nil {
		return err
	}
	return nil
}
