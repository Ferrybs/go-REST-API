package controllers

import (
	"blog/api/src/domain"
	"blog/api/src/usecase"
	"blog/api/src/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostController struct {
	PostUsecase usecase.PostUsecase
}

func NewPostController(PostUsecase *usecase.PostUsecase) *PostController {
	return &PostController{PostUsecase: *PostUsecase}
}

func (b *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var PostDTO domain.PostDTO
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.JsonWriter(w, http.StatusBadRequest, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	err = json.NewDecoder(r.Body).Decode(&PostDTO)
	if err != nil {
		utils.JsonWriter(w, http.StatusBadRequest, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	PostDTO.UserID = userID
	Post, err := b.PostUsecase.CreatePost(&PostDTO)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "Post post created successfully",
		Data:    Post,
	})
}

func (b *PostController) FindAllPost(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	Post, err := b.PostUsecase.FindAllPost(page, limit)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
	}
	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "Post posts fetched successfully",
		Data:    Post,
	})
}

func (b *PostController) FindAllPostByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	Post, err := b.PostUsecase.FindAllPostByUsername(username, page, limit)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "Post posts fetched successfully",
		Data:    Post,
	})
}

func (b *PostController) FindPostByPostID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	Post, err := b.PostUsecase.FindPostByID(domain.PostID(postID))
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	if Post == nil {
		utils.JsonWriter(w, http.StatusNotFound, domain.ResponseJson{
			Ok:      false,
			Message: "Post not found",
		})
		return
	}
	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "Post post fetched successfully",
		Data:    Post,
	})
}

func (b *PostController) DeletePostByPostID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.JsonWriter(w, http.StatusBadRequest, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	err = b.PostUsecase.DeletePostByPostID(userID, domain.PostID(postID))
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "Post post deleted successfully",
	})
}
