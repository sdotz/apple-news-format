package components

type figureStruct struct {
	componentStruct
	URL                  string `json:"URL"`
	AccessibilityCaption string `json:"accessibilityCaption,omitempty"`
	//Additions []*ComponentLink `json:"additions"`
	Caption         string `json:"caption,omitempty"`
	ExplicitContent bool   `json:"explicitContent,omitempty"`
}

func NewFigure() *figureStruct {
	i := figureStruct{}
	i.SetRole("figure")
	return &i
}
