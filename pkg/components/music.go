package components


func NewMusic() *audioStruct{
	i := audioStruct{}
	i.SetRole("music")
	return &i
}
