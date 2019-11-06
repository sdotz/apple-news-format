package components

type VideoStruct struct {
	componentStruct
	URL                  string  `json:"URL"`
	AccessibilityCaption string  `json:"accessibilityCaption,omitempty"`
	AspectRatio          float64 `json:"aspectRatio,omitempty"`
	Caption              string  `json:"caption,omitempty"`
	ExplicitContent      bool    `json:"explicitContent,omitempty"`
	StillURL             string  `json:"stillURL,omitempty"`
}

func NewVideo() *VideoStruct {
	i := VideoStruct{}
	i.SetRole("video")
	return &i
}
