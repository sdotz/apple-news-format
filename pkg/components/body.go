package components

type BodyStruct struct {
	TextStruct
}

func NewBody() *BodyStruct {
	b := BodyStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	b.SetRole("body")
	b.SetLayout("default-body")
	return &b
}
