package web

type ProductResponse struct {
	IdProduct int    `json:"idProduct"`
	Qty       int    `json:"qty"`
	Name      string `json:"name"`
}
