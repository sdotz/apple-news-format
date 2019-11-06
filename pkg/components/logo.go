package components

type LogoStruct struct {
	componentStruct
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	//Additions []*ComponentLink `json:"additions"`
	Caption         string `json:"caption,omitempty"`
	ExplicitContent bool   `json:"explicitContent,omitempty"`
}

func NewLogo() *LogoStruct {
	i := LogoStruct{}
	i.SetRole("logo")
	return &i
}
