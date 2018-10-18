package components

type introStruct struct {
	textStruct
}

func NewIntro() *introStruct {
	i := introStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("intro")
	return &i
}