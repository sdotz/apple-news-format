package styles

type ComponentTextStyle struct {
	*TextStyle
	DropCapStyle       *DropCapStyle `json:"dropCapStyle,omitempty"`
	FirstLineIndent    int           `json:"firstLineIndent,omitempty"`
	HangingPunctuation bool          `json:"hangingPunctuation,omitempty"`
	Hyphenation        bool          `json:"hyphenation,omitempty"`
	LineHeight         int           `json:"lineHeight,omitempty"`
	LinkStyle          *TextStyle    `json:"linkStyle,omitempty"`
	TextAlignment      string        `json:"textAlignment,omitempty"`
}
