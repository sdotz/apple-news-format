package styles

type Shadow struct {
	Color   Color   `json:"color"`
	Offset  Offset  `json:"offset"`
	Opacity float64 `json:"opacity"`
	Radius  float64 `json:"radius"`
}

type Offset struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
