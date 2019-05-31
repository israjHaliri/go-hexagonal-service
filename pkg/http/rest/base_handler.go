package rest

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
