package types

type EchoResponse struct {
	Args    map[string]interface{} `json:"args"`
	Data    map[string]interface{} `json:"data"`
	Headers map[string]string      `json:"headers"`
	URL     string                 `json:"url"`
}
