package components

type IllustratorStruct struct {
	TextStruct
}

func NewIllustrator() *IllustratorStruct {
	i := IllustratorStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("illustrator")
	return &i
}
