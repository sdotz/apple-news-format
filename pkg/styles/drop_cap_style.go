package styles

type DropCapStyle struct {
	BackgroundColor     *Color  `json:"backgroundColor,omitempty"`
	FontName            string `json:"fontName,omitempty"`
	NumberOfCharacters  int    `json:"numberOfCharacters,omitempty"`
	NumberOfLines       int    `json:"numberOfLines,omitempty"`
	NumberOfRaisedLines int    `json:"numberOfRaisedLines,omitempty"`
	Padding             int    `json:"padding,omitempty"`
	TextColor           *Color  `json:"textColor,omitempty"`
}
