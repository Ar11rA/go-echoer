package types

type EchoResponse struct {
	Args    map[string]interface{} `json:"args"`
	Data    map[string]interface{} `json:"data"`
	Headers map[string]string      `json:"headers"`
	URL     string                 `json:"url"`
}

type QuoteResponse struct {
	ID           string   `json:"_id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}
