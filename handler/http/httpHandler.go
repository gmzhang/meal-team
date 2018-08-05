package http

import (
	"github.com/gmzhang/meal-team/model/usecase"
	"github.com/labstack/echo"
	"time"
	"net/http"
)

type HttpHandler struct {
	ucase usecase.MealTeamUsecase
}

func NewHttpHandler(e *echo.Echo, ucase usecase.MealTeamUsecase) *HttpHandler {
	handler := HttpHandler{
		ucase: ucase,
	}
	e.GET("/", handler.index)
	home := e.Group("/meal_team")
	apiGroup := home.Group("/api")
	apiGroup.GET("/create_meal_team", handler.CreateMealTeam)

	return &handler
}

func (h *HttpHandler) index(c echo.Context) error {
	loc, _ := time.LoadLocation("Local")

	return c.String(http.StatusOK, "This is list-live micro server.("+time.Now().In(loc).String()+")")
}
func (h *HttpHandler) CreateMealTeam(c echo.Context) error {
	loc, _ := time.LoadLocation("Local")

	return c.String(http.StatusOK, "This is list-live micro server.("+time.Now().In(loc).String()+")")
}
