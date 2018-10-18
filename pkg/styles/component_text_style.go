package styles

type ComponentTextStyle struct {
	BackgroundColor    *Color        `json:"backgroundColor,omitempty"`
	DropCapStyle       *DropCapStyle `json:"dropCapStyle,omitempty"`
	FirstLineIndent    int           `json:"firstLineIndent,omitempty"`
	FontFamily         string        `json:"fontFamily,omitempty"`
	FontName           string        `json:"fontName,omitempty"`
	FontSize           int           `json:"fontSize,omitempty"`
	FontStyle          *FontStyle    `json:"fontStyle,omitempty"`
	FontWeight         string        `json:"fontWeight,omitempty"`
	FontWidth          string        `json:"fontWidth,omitempty"`
	HangingPunctuation bool          `json:"hangingPunctuation,omitempty"`
	Hyphenation        bool          `json:"hyphenation,omitempty"`
	LineHeight         int           `json:"lineHeight,omitempty"`
	LinkStyle          *TextStyle    `json:"linkStyle,omitempty"`
}
