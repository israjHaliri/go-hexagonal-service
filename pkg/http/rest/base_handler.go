package rest

const SecretJWT = "MYSECRETTOCHANG3"

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
