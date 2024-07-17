package presentation

import (
	dto "go-api/app/domain/dtos"
	service "go-api/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type eventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) eventController {
	return eventController{
		eventService,
	}
}

func (e *eventController) CreateEvent(ctx *gin.Context) {
	var request dto.CreateEventRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	savedEvent, err := e.eventService.CreateEvent(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, savedEvent)
}

func (e *eventController) GetEvents(ctx *gin.Context) {
	events := e.eventService.GetEvents()
	ctx.JSON(http.StatusOK, events)
}

func (e *eventController) GetEventById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := Response{
			Message: "event id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	eventId, err := strconv.Atoi(id)
	if err != nil {
		response := Response{
			Message: "event id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	event, err := e.eventService.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}
