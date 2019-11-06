package components

type AudioStruct struct {
	componentStruct
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	Caption              string `json:"caption,omitempty"`
	ExplicitContent      bool   `json:"explicitContent,omitempty"`
	ImageURL             string `json:"imageURL,omitempty"`
}

func NewAudio() *AudioStruct {
	i := AudioStruct{}
	i.SetRole("audio")
	return &i
}
