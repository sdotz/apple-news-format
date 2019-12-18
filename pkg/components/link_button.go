package components

import "github.com/sdotz/apple-news-format/pkg/styles"

type LinkButtonStruct struct {
	componentStruct
	URL       string                     `json:"URL"`
	Text      string                     `json:"text"`
	TextStyle *styles.ComponentTextStyle `json:"textStyle"`
}

func NewLinkButton() *LinkButtonStruct {
	i := LinkButtonStruct{}
	i.SetRole("link_button")
	return &i
}
