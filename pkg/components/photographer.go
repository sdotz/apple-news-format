package components

type photographerStruct struct {
	textStruct
}

func NewPhotographer() *photographerStruct{
	i := photographerStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("photographer")
	return &i
}
