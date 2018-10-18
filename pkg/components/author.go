package components

type authorStruct struct {
	textStruct
}

func NewAuthor() *authorStruct {
	i := authorStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("author")
	return &i
}
