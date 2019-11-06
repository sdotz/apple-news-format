package components

type IntroStruct struct {
	TextStruct
}

func NewIntro() *IntroStruct {
	i := IntroStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("intro")
	return &i
}
