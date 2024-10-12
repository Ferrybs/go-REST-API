package domain

type ResponseJson struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}
