package response

var EmptyData struct {}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Meta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}