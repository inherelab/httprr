package app

// BodyModel data reply model
type BodyModel struct {
	Url string
	Origin string
	Body string
	Json interface{}
	Args map[string]string
	Form map[string]string
	Files map[string]string
	Headers map[string]string
}

// QueryModel data reply model
type QueryModel struct {
	Url string
	Origin string
	Args map[string]string
	Headers map[string]string
}
