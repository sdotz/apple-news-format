package components

type illustratorStruct struct {
	textStruct
}

func NewIllustrator() *illustratorStruct{
	i := illustratorStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("illustrator")
	return &i
}
