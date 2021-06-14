package user

type LoginOutputFormat struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
