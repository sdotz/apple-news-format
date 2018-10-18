package components

type captionStruct struct {
	textStruct
}

func NewCaption() *captionStruct {
	i := captionStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("caption")
	return &i
}