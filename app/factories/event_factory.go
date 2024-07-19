package factory

import (
	"go-api/app/infra"
	"go-api/app/presentation"
	service "go-api/app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEvent(server *gin.Engine, db *gorm.DB) {
	Token := infra.NewJwtAdapter()
	EventRepository := infra.NewEventRepository(db)
	SubscriptionRepository := infra.NewSubscriptionRepository(db)

	EventService := service.NewEventService(EventRepository, SubscriptionRepository)
	SubscriptionService := service.NewSubscriptionService(SubscriptionRepository)

	EventController := presentation.NewEventController(EventService, SubscriptionService)

	server.GET("/event/:id", EventController.GetEventById)
	server.GET("/events", EventController.GetEvents)

	AuthUserMiddleware := presentation.NewAuthMiddleware(Token, "USER")
	protectedUsersRouter := server.Group("")
	protectedUsersRouter.Use(AuthUserMiddleware.Verify)
	protectedUsersRouter.POST("/event/:id/subscribe", EventController.Subscribe)
	protectedUsersRouter.DELETE("/event/:id/unsubscribe", EventController.Unsubscribe)

	AuthAdminMiddleware := presentation.NewAuthMiddleware(Token, "ADMIN")
	protectedAdminsRouter := server.Group("")
	protectedAdminsRouter.Use(AuthAdminMiddleware.Verify)
	protectedAdminsRouter.GET("/event/:id/users", EventController.GetEventUsers)
	protectedAdminsRouter.POST("/event", EventController.CreateEvent)
}
