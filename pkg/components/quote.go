package components

type quoteStruct struct {
	textStruct
}

func NewQuote() *quoteStruct {
	i := quoteStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("quote")
	return &i
}
