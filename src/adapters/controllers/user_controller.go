package controllers

import (
	"blog/api/src/domain"
	"blog/api/src/usecase"
	"blog/api/src/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: *userUsecase}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO domain.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		utils.JsonWriter(w, http.StatusBadRequest, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	user, err := c.userUsecase.CreateUser(&userDTO)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return

	}

	utils.JsonWriter(w, http.StatusCreated, domain.ResponseJson{
		Ok:      true,
		Message: "User created successfully",
		Data:    user,
	})
}

func (c *UserController) FindByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		utils.JsonWriter(w, http.StatusBadRequest, domain.ResponseJson{
			Ok:      false,
			Message: "Username is required",
		})
		return

	}

	user, err := c.userUsecase.FindByUsername(username)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "User found",
		Data:    user,
	})
}

func (c *UserController) GetUserAccessToken(w http.ResponseWriter, r *http.Request) {
	var userLoginDTO domain.UserLoginDTO

	err := json.NewDecoder(r.Body).Decode(&userLoginDTO)
	if err != nil {
		utils.JsonWriter(w, http.StatusBadRequest, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	userLoginView, err := c.userUsecase.GetUserAccessToken(&userLoginDTO)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, domain.ResponseJson{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	utils.JsonWriter(w, http.StatusOK, domain.ResponseJson{
		Ok:      true,
		Message: "User logged in successfully",
		Data:    userLoginView,
	})
}
