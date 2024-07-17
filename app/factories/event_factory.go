package factory

import (
	domain "go-api/app/domain/models"
	"go-api/app/infra"
	"go-api/app/presentation"
	service "go-api/app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEvent(server *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&domain.Event{})

	EventRepository := infra.NewEventRepository(db)
	EventService := service.NewEventService(EventRepository)
	EventController := presentation.NewEventController(EventService)

	server.GET("/event/:id", EventController.GetEventById)
	server.GET("/events", EventController.GetEvents)
	server.POST("/event", EventController.CreateEvent)
}
