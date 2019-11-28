package components

import (
	"github.com/sdotz/apple-news-format/pkg/styles"
)

type DividerStruct struct {
	componentStruct
	Stroke styles.StrokeStyle `json:"stroke"`
}

func NewDivider() *DividerStruct {
	c := DividerStruct{}
	c.SetRole("divider")
	return &c
}
