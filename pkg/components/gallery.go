package components

type galleryStruct struct {
	componentStruct
	Items []galleryItem `json:"items"`
}

type galleryItem struct {
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	Caption              string `json:"caption,omitempty"`
	ExplicitContent      bool   `json:"explicitContent,omitempty"`
}

func NewGallery() *galleryStruct {
	i := galleryStruct{}
	i.SetRole("gallery")
	return &i
}
