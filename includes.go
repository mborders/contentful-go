package contentful

type Includes struct {
	Asset []IncludesAsset `json:"Asset"`
}

type IncludesAsset struct {
	Metadata struct {
		Tags []interface{} `json:"tags"`
	} `json:"metadata"`
	Sys struct {
		Space struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"space"`
		ID          string `json:"id"`
		Type        string `json:"type"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		Environment struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"environment"`
		Revision int    `json:"revision"`
		Locale   string `json:"locale"`
	} `json:"sys"`
	Fields struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		File        struct {
			URL         string               `json:"url"`
			Details     IncludesAssetDetails `json:"details"`
			FileName    string               `json:"fileName"`
			ContentType string               `json:"contentType"`
		} `json:"file"`
	} `json:"fields"`
}

type IncludesAssetDetails struct {
	Size  int `json:"size"`
	Image struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"image"`
}
