package factory

import (
	"go-api/app/infra"
	"go-api/app/presentation"
	service "go-api/app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUser(server *gin.Engine, db *gorm.DB) {
	Cryptography := infra.BcryptAdapter{}
	Token := infra.NewJwtAdapter()
	UserRepository := infra.NewUserRepository(db)

	UserService := service.NewUserService(UserRepository, Cryptography, Token)

	AuthMiddleware := presentation.NewAuthMiddleware(Token, "USER")
	UserController := presentation.NewUserController(UserService)

	server.POST("/signin", UserController.SignIn)
	server.POST("/signup", UserController.SignUp)

	protectedRouter := server.Group("")
	protectedRouter.Use(AuthMiddleware.Verify)
	protectedRouter.PUT("/users/me", UserController.Update)
	protectedRouter.GET("/users/me", UserController.GetUser)
}
