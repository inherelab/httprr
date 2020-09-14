package app

// BodyModel data reply model
type BodyModel struct {
	Url string `json:"url"`
	Origin string `json:"origin"`
	Method string `json:"method"`
	Body string `json:"body"`
	Json interface{} `json:"json"`
	Args map[string]string `json:"args"`
	Form map[string]string `json:"form"`
	Files map[string]string `json:"files"`
	Headers map[string]string `json:"headers"`
}

// QueryModel data reply model
type QueryModel struct {
	Url string `json:"url"`
	Origin string `json:"origin"`
	Method string `json:"method"`
	Args map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
}
