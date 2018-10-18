package components

type embedWebVideoStruct struct {
	componentStruct
	URL                  string  `json:"URL"`
	AccessibilityCaption string  `json:"accessibilityCaption,omitempty"`
	AspectRatio          float64 `json:"aspectRatio,omitempty"`
	Caption              string  `json:"caption,omitempty"`
	ExplicitContent      bool    `json:"explicitContent,omitempty"`
}

func NewEmbedWebVideo() *embedWebVideoStruct {
	i := embedWebVideoStruct{}
	i.SetRole("embedwebvideo")
	return &i
}
