package routes

import (
	"blog/api/src/adapters/controllers"
	"blog/api/src/adapters/repositories"
	"blog/api/src/infrastructure/middlewares"
	"blog/api/src/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	userRepository := repositories.NewUserRepositoryMemory()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)
	r.Mount("/users", UserRoutes(userController))
	PostRepository := repositories.NewPostRepositoryMemory()
	PostUsecase := usecase.NewPostUsecase(PostRepository, userRepository)
	PostController := controllers.NewPostController(PostUsecase)
	r.Mount("/post", PostRoutes(PostController))
	return r
}

func PostRoutes(PostController *controllers.PostController) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{postID}", PostController.FindPostByPostID)
	r.Get("/", PostController.FindAllPost)
	r.Get("/user/{username}", PostController.FindAllPostByUsername)
	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWT)
		r.Post("/", PostController.CreatePost)
		r.Delete("/{postID}", PostController.DeletePostByPostID)
	})
	return r
}

func UserRoutes(userController *controllers.UserController) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", userController.CreateUser)
	r.Post("/login", userController.GetUserAccessToken)
	r.Get("/{username}", userController.FindByUsername)
	return r
}
