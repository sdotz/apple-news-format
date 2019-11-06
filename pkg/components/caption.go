package components

type CaptionStruct struct {
	TextStruct
}

func NewCaption() *CaptionStruct {
	i := CaptionStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("caption")
	return &i
}
