package components

func NewMusic() *AudioStruct {
	i := AudioStruct{}
	i.SetRole("music")
	return &i
}
