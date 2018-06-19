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
	SetFormat(Format)
	SetText(string)
	SetTextStyle(styles.ComponentTextStyle)
	SetInlineTextStyles([]styles.InlineTextStyle)
}

type textStruct struct {
	componentStruct
	Format           Format                    `json:"format"`
	Text             string                    `json:"text"`
	InlineTextStyles []*styles.InlineTextStyle  `json:"inlineTextStyles,omitempty"`
	TextStyle        *styles.ComponentTextStyle `json:"textStyle,omitempty"`
}

func (t *textStruct) SetRole(role string) Component {
	t.componentStruct.Role = role
	return t
}

func (t *textStruct) Role() string {
	return t.componentStruct.Role
}

func (t *textStruct) SetFormat(format Format) Component {
	t.Format = format
	return t
}

func (t *textStruct) SetComponentTextStyle(componentTextStyle styles.ComponentTextStyle) Component {
	t.TextStyle = &componentTextStyle
	return t
}

func (t *textStruct) SetInlineTextStyles(inlineTextStyles []*styles.InlineTextStyle) Component {
	t.InlineTextStyles = inlineTextStyles
	return t
}

func (t *textStruct) SetText(text string) Component {
	t.Text = text
	return t
}
