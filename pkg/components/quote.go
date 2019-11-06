package components

type quoteStruct struct {
	TextStruct
}

func NewQuote() *quoteStruct {
	i := quoteStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("quote")
	return &i
}
