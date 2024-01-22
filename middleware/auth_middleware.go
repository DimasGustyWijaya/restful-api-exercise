package middleware

import (
	"encoding/json"
	"net/http"
	"restful-api/app"
	"restful-api/helper"
	"restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

type UserLogin struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type AuthUser struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func userLogin(id int) UserLogin {
	db := app.GetConnection()
	usersLogin := UserLogin{}

	rows, err := db.Query("SELECT name from user where id = (?)", id)
	helper.PanicIfError(err)

	if rows.Next() {
		erro := rows.Scan(&usersLogin.Name)
		helper.PanicIfError(erro)
	}

	return usersLogin
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	id := 1
	user := userLogin(id)

	if user.Name == request.Header.Get("X-API-Auth") {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		encoder := json.NewEncoder(writer)
		errEncode := encoder.Encode(webResponse)
		helper.PanicIfError(errEncode)

	}
}
