package response

type Auth struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}
