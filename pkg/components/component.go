package components

import (
	"github.com/sdotz/apple-news-format/pkg/layouts"
	"github.com/sdotz/apple-news-format/pkg/styles"
)

type Component interface {
	SetRole(string) Component
	SetIdentifier(string) Component
	SetStyle(*styles.ComponentStyle) Component
}

type componentStruct struct {
	Role       string                   `json:"role"`
	Identifier string                   `json:"identifier,omitempty"`
	Layout     *layouts.ComponentLayout `json:"componentLayout,omitempty"`
	//Behavior Behavior
	//Animation
	//Anchor
	Style *styles.ComponentStyle `json:"style,omitempty"`
}

func (c *componentStruct) SetIdentifier(identifier string) Component {
	c.Identifier = identifier
	return c
}

func (c *componentStruct) SetRole(role string) Component {
	c.Role = role
	return c
}

func (c *componentStruct) SetStyle(style *styles.ComponentStyle) Component {
	c.Style = style
	return c
}
