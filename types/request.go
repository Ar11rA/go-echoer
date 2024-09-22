package types

type EchoRequest struct {
	Text string `json:"text"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
