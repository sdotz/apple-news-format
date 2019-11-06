package components

const (
	MapTypeStandard  MapType = "standard"
	MapTypeHybrid    MapType = "hybrid"
	MapTypeSatellite MapType = "satellite"
)

type MapType string

type MapStruct struct {
	componentStruct
	Caption              string    `json:"caption,omitempty"`
	AccessibilityCaption string    `json:"accessibilityCaption,omitempty"`
	Items                []MapItem `json:"items,omitempty"`
	Latitude             float64   `json:"latitude"`
	Longitude            float64   `json:"longitude"`
	MapType              MapType   `json:"mapType,omitempty"`
	Span                 MapSpan   `json:"mapSpan,omitempty"`
}

type MapItem struct {
	Caption   string  `json:"mapType,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MapSpan struct {
	LatitudeDelta  float64 `json:"latitudeDelta"`
	LongitudeDelta float64 `json:"longitudeDelta"`
}

func newMap() *MapStruct {
	m := MapStruct{}
	m.SetRole("map")
	return &m
}
