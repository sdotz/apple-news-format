package components


type placeStruct struct {
	componentStruct
	Caption              string    `json:"caption,omitempty"`
	AccessibilityCaption string    `json:"accessibilityCaption,omitempty"`
	Latitude             float64   `json:"latitude"`
	Longitude            float64   `json:"longitude"`
	MapType              MapType   `json:"mapType,omitempty"`
	Span                 MapSpan   `json:"mapSpan,omitempty"`
}

func newPlace() *placeStruct {
	m := placeStruct{}
	m.SetRole("place")
	return &m
}

