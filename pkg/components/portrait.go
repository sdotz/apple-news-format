package components

type portraitStruct struct {
	componentStruct
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	//Additions []*ComponentLink `json:"additions"`
	Caption         string `json:"caption,omitempty"`
	ExplicitContent bool   `json:"explicitContent,omitempty"`
}

func NewPortrait() *portraitStruct {
	i := portraitStruct{}
	i.SetRole("portrait")
	return &i
}
