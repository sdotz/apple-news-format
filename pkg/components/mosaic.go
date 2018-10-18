package components

type mosaicStruct struct {
	componentStruct
	Items []galleryItem `json:"items"`
}

func NewMosaic() *mosaicStruct {
	i := mosaicStruct{}
	i.SetRole("mosaic")
	return &i
}
