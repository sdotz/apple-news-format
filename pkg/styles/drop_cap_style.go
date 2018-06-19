package styles

type DropCapStyle struct {
	BackgroundColor     Color  `json:"backgroundColor"`
	FontName            string `json:"fontName"`
	NumberOfCharacters  int    `json:"numberOfCharacters"`
	NumberOfLines       int    `json:"numberOfLines"`
	NumberOfRaisedLines int    `json:"numberOfRaisedLines"`
	Padding             int    `json:"padding"`
	TextColor           Color  `json:"textColor"`
}
