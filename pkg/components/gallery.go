package components

type GalleryStruct struct {
	componentStruct
	Items []GalleryItem `json:"items"`
}

type GalleryItem struct {
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	Caption              string `json:"caption,omitempty"`
	ExplicitContent      bool   `json:"explicitContent,omitempty"`
}

func NewGallery() *GalleryStruct {
	i := GalleryStruct{}
	i.SetRole("gallery")
	return &i
}
