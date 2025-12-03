package museums

// Struct for sending list of museums to frontend
type Museum struct {
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	APIBase string `json:"apiBase"`
}
