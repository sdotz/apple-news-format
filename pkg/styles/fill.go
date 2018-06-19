package styles

const (
	LinearGradientFillType = "linear_gradient"
)

type ColorStop struct {
	Color Color `json:"color"`
	Location int `json:"location,omitempty"`
}

type Fill interface {

}

type GradientFill struct {
	ColorStops []ColorStop `json:"colorStops"`
}

type LinearGradientFill struct {

	Type string
}


