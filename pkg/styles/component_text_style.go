package styles

type ComponentTextStyle struct {
	BackgroundColor    Color        `json:"backgroundColor"`
	DropCapStyle       DropCapStyle `json:"dropCapStyle"`
	FirstLineIndent    int          `json:"firstLineIndent"`
	FontFamily         string       `json:"fontFamily"`
	FontName           string       `json:"fontName"`
	FontSize           int          `json:"fontSize"`
	FontStyle          FontStyle    `json:"fontStyle"`
	FontWeight         string       `json:"fontWeight"`
	FontWidth          string       `json:"fontWidth"`
	HangingPunctuation bool         `json:"hangingPunctuation"`
	Hyphenation        bool         `json:"hyphenation"`
	LineHeight         int          `json:"lineHeight"`
}
