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

type textStruct struct {
	componentStruct
	Format           Format                     `json:"format,omitempty"`
	Text             string                     `json:"text"`
	InlineTextStyles []*styles.InlineTextStyle  `json:"inlineTextStyles,omitempty"`
	TextStyle        *styles.ComponentTextStyle `json:"textStyle,omitempty"`
}

func (t *textStruct) SetFormat(format Format) Text {
	t.Format = format
	return t
}

func (t *textStruct) SetTextStyle(componentTextStyle styles.ComponentTextStyle) Text {
	t.TextStyle = &componentTextStyle
	return t
}

func (t *textStruct) SetInlineTextStyles(inlineTextStyles []*styles.InlineTextStyle) Text {
	t.InlineTextStyles = inlineTextStyles
	return t
}

func (t *textStruct) SetText(text string) Text {
	t.Text = text
	return t
}
