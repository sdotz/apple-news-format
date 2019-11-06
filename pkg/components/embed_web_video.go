package components

type EmbedWebVideoStruct struct {
	componentStruct
	URL                  string  `json:"URL"`
	AccessibilityCaption string  `json:"accessibilityCaption,omitempty"`
	AspectRatio          float64 `json:"aspectRatio,omitempty"`
	Caption              string  `json:"caption,omitempty"`
	ExplicitContent      bool    `json:"explicitContent,omitempty"`
}

func NewEmbedWebVideo() *EmbedWebVideoStruct {
	i := EmbedWebVideoStruct{}
	i.SetRole("embedwebvideo")
	return &i
}
