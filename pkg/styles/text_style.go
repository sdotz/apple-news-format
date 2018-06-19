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
	BackgroundColor    Color             `json:"backgroundColor"`
	FontFamily         string            `json:"fontFamily"`
	FontName           string            `json:"fontName"`
	FontSize           int               `json:"fontSize"`
	FontStyle          FontStyle         `json:"fontStyle"`
	FontWeight         string            `json:"fontWeight"`
	FontWidth          string            `json:"fontWidth"`
	OrderedListItems   ListItemStyle     `json:"orderedListItems"`
	Strikethrough      TextDecoration    `json:"strikethrough"`
	Stroke             TextStrokeStyle   `json:"textStrokeStyle"`
	TextColor          Color             `json:"textColor"`
	TextShadow         Shadow            `json:"textShadow"`
	Tracking           float64           `json:"tracking"`
	Underline          TextDecoration    `json:"underline"`
	UnorderedListItems ListItemStyle     `json:"unorderedListItems"`
	VerticalAlignment  VerticalAlignment `json:"verticalAlignment"`
}
