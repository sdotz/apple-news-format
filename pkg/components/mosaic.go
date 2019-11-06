package components

type MosaicStruct struct {
	componentStruct
	Items []GalleryItem `json:"items"`
}

func NewMosaic() *MosaicStruct {
	i := MosaicStruct{}
	i.SetRole("mosaic")
	return &i
}
