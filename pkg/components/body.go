package components


type Body interface {
	Text
	Component
}

type bodyStruct struct {
	textStruct
}

func NewBody() *bodyStruct {
	b := bodyStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	b.SetRole("body")
	return &b
}


func (b *bodyStruct) SetRole(role string) Component {
	b.textStruct.SetRole(role)
	return b
}

func (b *bodyStruct) Role() string {
	return b.textStruct.Role()
}
