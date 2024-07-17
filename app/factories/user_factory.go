package factory

import (
	domain "go-api/app/domain/models"
	"go-api/app/infra"
	"go-api/app/presentation"
	service "go-api/app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUser(server *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&domain.User{})

	Cryptography := infra.BcryptAdapter{}
	Token := infra.NewJwtAdapter("jwt-secret")
	UserRepository := infra.NewUserRepository(db)

	UserService := service.NewUserService(UserRepository, Cryptography, Token)

	UserController := presentation.NewUserController(UserService)

	server.POST("/signin", UserController.SignIn)
	server.POST("/signup", UserController.SignUp)
	server.PUT("/user/:id", UserController.Update)
}
