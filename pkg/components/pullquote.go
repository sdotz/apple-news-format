package components

type pullQuoteStruct struct {
	textStruct
}

func NewPullQuote() *pullQuoteStruct {
	i := pullQuoteStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	i.SetRole("pullquote")
	return &i
}
