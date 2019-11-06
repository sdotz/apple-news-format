package components

type BylineStruct struct {
	TextStruct
}

func NewByline() *BylineStruct {
	i := BylineStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("byline")
	return &i
}
