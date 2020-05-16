package styles

type Color string

type StrokeStyleOption string

const (
	StrokeStyleValueSolid  StrokeStyleOption = "solid"
	StrokeStyleValueDashed StrokeStyleOption = "dashed"
	StrokeStyleValueDotted StrokeStyleOption = "dotted"
)

type Mask struct {
	Radius int    `json:"radius,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Border struct {
	All    StrokeStyle `json:"all,omitempty"`
	Bottom bool        `json:"bottom"`
	Left   bool        `json:"left"`
	Right  bool        `json:"right"`
	Top    bool        `json:"top"`
}

type StrokeStyle struct {
	Color Color             `json:"color,omitempty"`
	Style StrokeStyleOption `json:"style,omitempty"`
	Width string            `json:"width"`
}

type ComponentStyle struct {
	BackgroundColor Color   `json:"backgroundColor,omitempty"`
	Border          *Border `json:"border,omitempty"`
	Fill            *Fill   `json:"fill,omitempty"`
	Mask            *Mask   `json:"mask,omitempty"`
}
