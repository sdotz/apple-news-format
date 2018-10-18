package components

type imageStruct struct {
	componentStruct
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	//Additions []*ComponentLink `json:"additions"`
	Caption         string `json:"caption,omitempty"`
	ExplicitContent bool   `json:"explicitContent,omitempty"`
}

func NewImage() *imageStruct {
	i := imageStruct{}
	i.SetRole("image")
	return &i
}
