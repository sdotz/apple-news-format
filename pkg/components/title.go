package components

type titleStruct struct {
	textStruct
}

func NewTitle() *titleStruct {
	t := titleStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	t.SetRole("title")
	return &t
}