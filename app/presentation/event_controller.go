package presentation

import (
	dto "go-api/app/domain/dtos"
	service "go-api/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type eventController struct {
	eventService        service.EventService
	subscriptionService service.SubscriptionService
}

func NewEventController(eventService service.EventService, subscriptionService service.SubscriptionService) eventController {
	return eventController{
		eventService,
		subscriptionService,
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

	eventId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := Response{
			Message: "event id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	event, err := e.eventService.GetEventById(uint(eventId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (e *eventController) GetEventUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := Response{
			Message: "event id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	eventId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := Response{
			Message: "event id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	users, err := e.eventService.GetEventUsers(uint(eventId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (e *eventController) Subscribe(ctx *gin.Context) {
	userId := ctx.GetUint("x-user-id")

	eventId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response := Response{
			Message: "event id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = e.subscriptionService.Subscribe(userId, uint(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (e *eventController) Unsubscribe(ctx *gin.Context) {
	userId := ctx.GetUint("x-user-id")

	eventId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response := Response{
			Message: "event id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = e.subscriptionService.Unsubscribe(userId, uint(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
