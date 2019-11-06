package components

type pullQuoteStruct struct {
	TextStruct
}

func NewPullQuote() *pullQuoteStruct {
	i := pullQuoteStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("pullquote")
	return &i
}
