package web

type ProductUpdateRequest struct {
	IdProduct int    `validate:"required"`
	Name      string `validate:"required"`
	Qty       int    `validate:"required"`
}
