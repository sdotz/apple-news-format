package components

const (
	MapTypeStandard  MapType = "standard"
	MapTypeHybrid    MapType = "hybrid"
	MapTypeSatellite MapType = "satellite"
)

type MapType string

type mapStruct struct {
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

func newMap() *mapStruct {
	m := mapStruct{}
	m.SetRole("map")
	return &m
}
