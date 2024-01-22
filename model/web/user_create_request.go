package web

type UserCreateRequest struct {
	Name string `validate:"required,min=1,max=100"`
}
