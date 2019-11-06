package components

type TitleStruct struct {
	TextStruct
}

func NewTitle() *TitleStruct {
	t := TitleStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	t.SetRole("title")
	return &t
}
