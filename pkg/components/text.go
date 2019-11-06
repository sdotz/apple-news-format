package components

import (
	"github.com/sdotz/apple-news-format/pkg/styles"
)

type Format string

const (
	FormatNone     = "none"
	FormatMarkdown = "markdown"
	FormatHtml     = "html"
)

type Text interface {
	Component
	SetFormat(Format) Text
	SetText(string) Text
	SetTextStyle(styles.ComponentTextStyle) Text
	SetInlineTextStyles([]*styles.InlineTextStyle) Text
}

type TextStruct struct {
	componentStruct
	Format           Format                     `json:"format,omitempty"`
	Text             string                     `json:"text"`
	InlineTextStyles []*styles.InlineTextStyle  `json:"inlineTextStyles,omitempty"`
	TextStyle        *styles.ComponentTextStyle `json:"textStyle,omitempty"`
}

func (t *TextStruct) SetLayout(layout interface{}) error {
	return t.componentStruct.SetLayout(layout)
}

func (t *TextStruct) SetFormat(format Format) Text {
	t.Format = format
	return t
}

func (t *TextStruct) SetTextStyle(componentTextStyle styles.ComponentTextStyle) Text {
	t.TextStyle = &componentTextStyle
	return t
}

func (t *TextStruct) SetInlineTextStyles(inlineTextStyles []*styles.InlineTextStyle) Text {
	t.InlineTextStyles = inlineTextStyles
	return t
}

func (t *TextStruct) SetText(text string) Text {
	t.Text = text
	return t
}
