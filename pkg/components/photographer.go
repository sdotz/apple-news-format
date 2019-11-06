package components

type photographerStruct struct {
	TextStruct
}

func NewPhotographer() *photographerStruct {
	i := photographerStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("photographer")
	return &i
}
