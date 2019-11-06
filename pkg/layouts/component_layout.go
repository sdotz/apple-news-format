package layouts

import (
	"github.com/sdotz/apple-news-format/pkg/properties"
)

type HorizontalContentAlignmentOption string
type IgnoreDocumentGutterOption string
type IgnoreDocumentMarginOption string

const (
	HorizontalAlignmentValueLeft   HorizontalContentAlignmentOption = "left"
	HorizontalAlignmentValueCenter HorizontalContentAlignmentOption = "center"
	HorizontalAlignmentValueRight  HorizontalContentAlignmentOption = "right"

	IgnoreDocumentGutterValueTrue  IgnoreDocumentGutterOption = "true"
	IgnoreDocumentGutterValueFalse IgnoreDocumentGutterOption = "false"
	IgnoreDocumentGutterValueNone  IgnoreDocumentGutterOption = "none"
	IgnoreDocumentGutterValueLeft  IgnoreDocumentGutterOption = "left"
	IgnoreDocumentGutterValueRight IgnoreDocumentGutterOption = "right"
	IgnoreDocumentGutterValueBoth  IgnoreDocumentGutterOption = "both"

	IgnoreDocumentMarginValueTrue  IgnoreDocumentMarginOption = "true"
	IgnoreDocumentMarginValueFalse IgnoreDocumentMarginOption = "false"
	IgnoreDocumentMarginValueNone  IgnoreDocumentMarginOption = "none"
	IgnoreDocumentMarginValueLeft  IgnoreDocumentMarginOption = "left"
	IgnoreDocumentMarginValueRight IgnoreDocumentMarginOption = "right"
	IgnoreDocumentMarginValueBoth  IgnoreDocumentMarginOption = "both"
)

type ComponentLayout struct {
	ColumnSpan                 int                              `json:"columnSpan"`
	ColumnStart                int                              `json:"columnStart"`
	ContentInset               ContentInset                     `json:"contentInset,omitempty"`
	HorizontalContentAlignment HorizontalContentAlignmentOption `json:"horizontalContentAlignment,omitempty"`
	IgnoreDocumentGutter       IgnoreDocumentGutterOption       `json:"ignoreDocumentGutter,omitempty"`
	IgnoreDocumentMargin       IgnoreDocumentMarginOption       `json:"ignoreDocumentMargin,omitempty"`
	MaximumContentWidth        string                           `json:"maximumContentWidth,omitempty"`
	MinimumHeight              string                           `json:"minimumHeight,omitempty"`
	Margin                     properties.Margin                `json:"margin,omitempty"`
}

type ContentInset struct {
	Top    bool `json:"top,omitempty"`
	Left   bool `json:"left,omitempty"`
	Right  bool `json:"right,omitempty"`
	Bottom bool `json:"bottom,omitempty"`
}
