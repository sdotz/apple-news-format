package components

import (
	"fmt"

	"github.com/sdotz/apple-news-format/pkg/layouts"
	"github.com/sdotz/apple-news-format/pkg/styles"
)

type Component interface {
	SetRole(string) Component
	SetIdentifier(string) Component
	SetStyle(*styles.ComponentStyle) Component
	SetLayout(interface{}) error
}

type componentStruct struct {
	Role       string `json:"role"`
	Identifier string `json:"identifier,omitempty"`
	// Layout can be either a *layouts.ComponentLayout or a string referring to a layout
	Layout interface{} `json:"layout,omitempty"`
	//Behavior Behavior
	//Animation
	//Anchor
	Style *styles.ComponentStyle `json:"style,omitempty"`
}

func (c *componentStruct) SetLayout(layout interface{}) error {
	layoutString, isString := layout.(string)
	if isString {
		c.Layout = layoutString
	}

	layoutStruct, isStruct := layout.(*layouts.ComponentLayout)
	if isStruct {
		c.Layout = layoutStruct
	}

	if !isString && !isStruct {
		return fmt.Errorf("layout must be set to either a string or object of type *layouts.ComponentLayout")
	}

	return nil
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
