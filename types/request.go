package types

type EchoRequest struct {
	Text string `json:"text"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RedisDataRequest struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}
