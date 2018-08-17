package http

import (
	"github.com/gmzhang/meal-team/model/usecase"
	"github.com/gmzhang/meal-team/model"
	"github.com/labstack/echo"
	"time"
	"net/http"
	"strconv"
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
	apiGroup.POST("/create_meal_team", handler.CreateMealTeam)
	apiGroup.GET("/get_sys_meal_team_name", handler.GetALLMealTeamName)
	apiGroup.GET("/get_meal_team/:mealTeamId", handler.GetMealTeam)
	apiGroup.GET("/get_meal_team_member/:mealTeamId", handler.GetMealTeamMember)
	apiGroup.GET("/get_meal_team_restaurant/:mealTeamId", handler.GetMealTeamRestaurant)
	apiGroup.POST("/confirm_meal_team_restaurant", handler.UpMealTeamRestaurantId)
	apiGroup.POST("/broadcast_meal_team_member", handler.BroadcastMealTeamMember)
	apiGroup.POST("/turn_over_meal_team_manager", handler.TurnOverMealTeamer)
	apiGroup.POST("/join_meal_team", handler.JoinMealTeam)
	apiGroup.POST("/notify_meal_team_manager", handler.NotifyMealTeamManager)
	apiGroup.GET("/get_meal_team_manager_running/:mealTeamId", handler.GetMealTeamManagerRunning)
	apiGroup.GET("/get_meal_teams_by_openid/:openid", handler.GetMealTeamsByOpenid)

	return &handler
}

func (h *HttpHandler) index(c echo.Context) error {
	loc, _ := time.LoadLocation("Local")

	return c.String(http.StatusOK, "This is list-live micro server.("+time.Now().In(loc).String()+")")
}
func (h *HttpHandler) CreateMealTeam(c echo.Context) error {
	postInfo := struct {
		Name   string `json:"name"`
		Openid string `json:"openid"`
		Nick   string `json:"nick"`
		Avatar string `json:"avatar"`
	}{}

	if err := c.Bind(&postInfo); err != nil {
		return c.JSON(http.StatusOK, model.GetErrorMap(err))
	}
	mealTeam, err := h.ucase.Create(postInfo.Name, postInfo.Openid, postInfo.Nick, postInfo.Avatar)
	result := model.GetErrorMap(err)
	result["mealTeam"] = mealTeam
	return c.JSON(http.StatusOK, result)
}
func (h *HttpHandler) GetALLMealTeamName(c echo.Context) error {
	names, err := h.ucase.GetALLMealTeamName()

	result := model.GetErrorMap(err)

	result["names"] = names

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) GetMealTeam(c echo.Context) error {
	mealTeamIdStr := c.Param("mealTeamId")
	mealTeamId, err := strconv.Atoi(mealTeamIdStr)
	mealTeam, err := h.ucase.GetMealTeam(mealTeamId)

	result := model.GetErrorMap(err)

	result["mealTeam"] = mealTeam

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) GetMealTeamMember(c echo.Context) error {
	mealTeamIdStr := c.Param("mealTeamId")
	mealTeamId, err := strconv.Atoi(mealTeamIdStr)
	members, err := h.ucase.GetMealTeamMember(mealTeamId)

	result := model.GetErrorMap(err)

	result["members"] = members

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) GetMealTeamRestaurant(c echo.Context) error {
	mealTeamIdStr := c.Param("mealTeamId")
	mealTeamId, err := strconv.Atoi(mealTeamIdStr)
	restaurants, err := h.ucase.GetMealTeamALLRestaurant(mealTeamId)

	result := model.GetErrorMap(err)

	result["restaurants"] = restaurants

	return c.JSON(http.StatusOK, result)
}
func (h *HttpHandler) UpMealTeamRestaurantId(c echo.Context) error {
	postInfo := struct {
		MealTeamId   int `json:"mealTeamId"`
		RestaurantId int `json:"restaurantId"`
	}{}

	if err := c.Bind(&postInfo); err != nil {
		return c.JSON(http.StatusOK, model.GetErrorMap(err))
	}
	err := h.ucase.UpdateMealTeamRestaurantId(postInfo.MealTeamId, postInfo.RestaurantId)
	result := model.GetErrorMap(err)
	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) BroadcastMealTeamMember(c echo.Context) error {
	postInfo := struct {
		MealTeamId int `json:"mealTeamId"`
	}{}

	if err := c.Bind(&postInfo); err != nil {
		return c.JSON(http.StatusOK, model.GetErrorMap(err))
	}
	err := h.ucase.BroadcastMealTeamMember(postInfo.MealTeamId)
	result := model.GetErrorMap(err)
	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) TurnOverMealTeamer(c echo.Context) error {
	postInfo := struct {
		MealTeamId int `json:"mealTeamId"`
	}{}

	if err := c.Bind(&postInfo); err != nil {
		return c.JSON(http.StatusOK, model.GetErrorMap(err))
	}
	err := h.ucase.TurnOverMealTeamer(postInfo.MealTeamId)
	result := model.GetErrorMap(err)
	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) JoinMealTeam(c echo.Context) error {
	postInfo := struct {
		MealTeamId int    `json:"mealTeamId"`
		Openid     string `json:"openid"`
		Nick       string `json:"nick"`
		Avatar     string `json:"avatar"`
	}{}

	if err := c.Bind(&postInfo); err != nil {
		return c.JSON(http.StatusOK, model.GetErrorMap(err))
	}
	err := h.ucase.JoinMealTeam(postInfo.MealTeamId, postInfo.Openid, postInfo.Nick, postInfo.Avatar)
	result := model.GetErrorMap(err)
	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) NotifyMealTeamManager(c echo.Context) error {
	postInfo := struct {
		MealTeamId int    `json:"mealTeamId"`
		Openid     string `json:"openid"`
	}{}

	if err := c.Bind(&postInfo); err != nil {
		return c.JSON(http.StatusOK, model.GetErrorMap(err))
	}
	err := h.ucase.TurnOverMealTeamer(postInfo.MealTeamId)
	result := model.GetErrorMap(err)
	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) GetMealTeamManagerRunning(c echo.Context) error {
	mealTeamIdStr := c.Param("mealTeamId")
	mealTeamId, err := strconv.Atoi(mealTeamIdStr)
	running, err := h.ucase.GetMealTeamManagerRunning(mealTeamId)

	result := model.GetErrorMap(err)

	result["running"] = running

	return c.JSON(http.StatusOK, result)
}
func (h *HttpHandler) GetMealTeamsByOpenid(c echo.Context) error {
	openidStr := c.Param("openid")
	mealTeams, err := h.ucase.GetMealTeamManagersByMemberOpenid(openidStr)

	result := model.GetErrorMap(err)

	result["mealTeams"] = mealTeams

	return c.JSON(http.StatusOK, result)
}
