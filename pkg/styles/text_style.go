package styles

type FontStyle string
type VerticalAlignment string

const (
	FontStyleNormal  FontStyle = "normal"
	FontStyleItalic  FontStyle = "italic"
	FontStyleOblique FontStyle = "oblique"

	VerticalAlignmentSuperscript VerticalAlignment = "superscript"
	VerticalAlignmentSubscript   VerticalAlignment = "subscript"
	VerticalAlignmentBaseline    VerticalAlignment = "baseline"
)

type TextStyle struct {
	BackgroundColor    *Color             `json:"backgroundColor,omitempty"`
	FontFamily         string             `json:"fontFamily,omitempty"`
	FontName           string             `json:"fontName,omitempty"`
	FontSize           int                `json:"fontSize,omitempty"`
	FontStyle          *FontStyle         `json:"fontStyle,omitempty"`
	FontWeight         string             `json:"fontWeight,omitempty"`
	FontWidth          string             `json:"fontWidth,omitempty"`
	OrderedListItems   *ListItemStyle     `json:"orderedListItems,omitempty"`
	Strikethrough      *TextDecoration    `json:"strikethrough,omitempty"`
	Stroke             *TextStrokeStyle   `json:"textStrokeStyle,omitempty"`
	TextColor          *Color             `json:"textColor,omitempty"`
	TextShadow         *Shadow            `json:"textShadow,omitempty"`
	Tracking           float64            `json:"tracking,omitempty"`
	Underline          *TextDecoration    `json:"underline,omitempty"`
	UnorderedListItems *ListItemStyle     `json:"unorderedListItems,omitempty"`
	VerticalAlignment  *VerticalAlignment `json:"verticalAlignment,omitempty"`
}
