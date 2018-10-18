package components

type logoStruct struct {
	componentStruct
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	//Additions []*ComponentLink `json:"additions"`
	Caption         string `json:"caption,omitempty"`
	ExplicitContent bool   `json:"explicitContent,omitempty"`
}

func NewLogo() *logoStruct {
	i := logoStruct{}
	i.SetRole("logo")
	return &i
}
