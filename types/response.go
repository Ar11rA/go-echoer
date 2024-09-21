package types

type EchoResponse struct {
	Args struct {
		Query string `json:"query"`
	} `json:"args"`
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}
