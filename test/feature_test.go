package test

import (
	"database/sql"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/helper"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"
	"strings"
	"testing"
	"time"
)

func setupDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/tokopa")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

}

func setupRouter() http.Handler {
	db := setupDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validate)
	userControler := controller.NewUserController(userService)

	router := app.NewRouter(userControler)

	return middleware.NewAuthMiddleware(router)

}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name":"jiraiya"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:4400/api/user", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Auth", "jeffersoN")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Result().StatusCode)

}
