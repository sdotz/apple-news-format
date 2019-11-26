package components

type ContainerStruct struct {
	componentStruct
	Components []Component `json:"components"`
}

func NewContainer() *ContainerStruct {
	i := ContainerStruct{}
	i.SetRole("container")
	return &i
}

func (c *ContainerStruct) Add(component Component) {
	c.Components = append(c.Components, component)
}
