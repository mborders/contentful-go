package contentful

// Tag model
type Tag struct {
	Sys struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		LinkType string `json:"linkType"`
	} `json:"sys"`
}
