package types

type LogRequest struct {
	ApplicationID string `json:"applicationId"`
	Logs          string `json:"logs"`
}

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

type MessagePublishRequest struct {
	Content string `json:"content" validate:"required"`
}
