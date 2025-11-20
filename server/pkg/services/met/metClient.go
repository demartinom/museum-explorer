package met

func NewClient() *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1"}
}
