package main

import (
	"os"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/gmzhang/meal-team/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gmzhang/meal-team/model/repository"
	"github.com/gmzhang/meal-team/model/usecase"
	"github.com/gmzhang/meal-team/handler/http"
)

func main() {
	initLogger()
	dbCoon := initDB()
	repo := repository.NewMealTeamRepository(dbCoon)

	useCase := usecase.NewMealTeamUsecase(repo)

	webService(useCase)

}

func webService(useCase usecase.MealTeamUsecase) {

	e := echo.New()
	e.Use(middleware.Recover())

	dataPath := os.Getenv("DATA_PATH")
	accessLogFile, err := os.OpenFile(dataPath+"access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalln("Failed create access log file")
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: accessLogFile,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out},"user_agent":"${user_agent}"}` + "\n",
	}))

	http.NewHttpHandler(e, useCase)

	err = e.Start(":8000")
	if err != nil {
		logrus.Fatalf("webService failed to start: %v", err)
	}

}

func initLogger() {

	dataPath := os.Getenv("DATA_PATH")

	file, err := os.OpenFile(dataPath+"main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Fatalln("Failed to log to file, using default stderr")
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func initDB() *sqlx.DB {

	connStr := config.DatabaseUser + ":" + config.DatabasePass + "@tcp(" + config.DatabaseHost +
		":" + config.DatabasePort + ")/" + config.DatabaseName + "?loc=Local&parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci"

	logrus.Infoln("connect db ....", connStr)
	var err error
	DB, err := sqlx.Open("mysql", connStr)

	if err == nil {
		err = DB.Ping()
	}

	if err != nil {
		logrus.Fatalf("sqlx system db connect error %s : %#v", connStr, err)
	}
	return DB
}
