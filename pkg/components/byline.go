package components

type bylineStruct struct {
	textStruct
}

func NewByline() *bylineStruct {
	i := bylineStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("byline")
	return &i
}
