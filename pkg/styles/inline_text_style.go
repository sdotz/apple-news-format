package styles

type InlineTextStyle struct {
	RangeLength int       `json:"rangeLength"`
	RangeStart  int       `json:"rangeStart"`
	TextStyle   TextStyle `json:"textStyle"`
}


