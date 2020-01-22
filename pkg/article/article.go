package article

import (
	"encoding/json"
	"io/ioutil"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sdotz/apple-news-format/pkg/components"
	"github.com/sdotz/apple-news-format/pkg/layouts"
	"github.com/sdotz/apple-news-format/pkg/properties"
	"github.com/sdotz/apple-news-format/pkg/styles"
)

type CoverArtType string
type BannerType string

const (
	CoverArtTypeImage = "image"

	BannerTypeAny          BannerType = "any"
	BannerTypeStandard     BannerType = "standard"
	BannerTypeDoubleHeight BannerType = "double_height"
	BannerTypeLarge        BannerType = "large"

	FORMAT_VERSION_1_7 = "1.7"
	FORMAT_VERSION_1_9 = "1.9"
)

type Article struct {
	Version             string                               `json:"version"`
	Identifier          string                               `json:"identifier"`
	Language            string                               `json:"language"`
	Title               string                               `json:"title"`
	Subtitle            string                               `json:"subtitle"`
	Layout              Layout                               `json:"layout"`
	AdvertisingSettings AdvertisingSettings                  `json:"advertisingSettings"`
	Metadata            ArticleMetadata                      `json:"metadata"`
	Components          []components.Component               `json:"components"`
	DocumentStyle       styles.DocumentStyle                 `json:"documentStyle"`
	ComponentTextStyles map[string]styles.ComponentTextStyle `json:"componentTextStyles,omitempty"`
	ComponentStyles     map[string]styles.ComponentStyle     `json:"componentStyles,omitempty"`
	TextStyles          map[string]styles.TextStyle          `json:"textStyles,omitempty"`
	ComponentLayouts    map[string]layouts.ComponentLayout   `json:"componentLayouts,omitempty"`
}

type Layout struct {
	Columns int `json:"columns"`
	Width   int `json:"width"`
	Margin  int `json:"margin"`
	Gutter  int `json:"gutter"`
}

type CoverArt struct {
	Type                 CoverArtType `json:"type"`
	URL                  string       `json:"URL"`
	AccessibilityCaption string       `json:"accessibilityCaption,omitempty"`
}

type AdvertisingSettings struct {
	BannerType        string            `json:"bannerType,omitempty"`
	DistanceFromMedia string            `json:"distanceFromMedia,omitempty"` // Should this take a "Measurement" type to support px, em etc.?
	Frequency         int               `json:"frequency,omitempty"`
	Layout            AdvertisingLayout `json:"advertisingLayout,omitempty"`
}

type AdvertisingLayout struct {
	Margin properties.Margin `json:"margin,omitempty"`
}

type LinkedArticle struct {
	URL          string `json:"url"`
	Relationship string `json:"relationship"`
}

type ArticleMetadata struct {
	Authors             []string               `json:"authors,omitempty"`
	CampaignData        map[string]interface{} `json:"campaignData,omitempty"`
	CanonicalURL        string                 `json:"canonicalURL,omitempty"`
	CoverArts           []CoverArt             `json:"coverArt,omitempty"`
	DateCreated         string                 `json:"dateCreated,omitempty"`
	DateModified        string                 `json:"dateModified,omitempty"`
	DatePublished       string                 `json:"datePublished,omitempty"`
	Excerpt             string                 `json:"excerpt,omitempty"`
	GeneratorIdentifier string                 `json:"generatorIdentifier,omitempty"`
	GeneratorName       string                 `json:"generatorName,omitempty"`
	GeneratorVersion    string                 `json:"generatorVersion,omitempty"`
	Keywords            []string               `json:"keywords,omitempty"`
	Links               []LinkedArticle        `json:"links,omitempty"`
	ThumbnailURL        string                 `json:"thumbnailURL,omitempty"`
	TransparentToolbar  string                 `json:"transparentToolbar,omitempty"`
	VideoURL            string                 `json:"videoURL,omitempty"`
}

func NewArticleWithDefaults() Article {
	identifier, _ := uuid.NewV4()
	a := Article{
		Language: "en",
		Version:  FORMAT_VERSION_1_9,
		Layout: Layout{
			Columns: 10,
			Gutter:  20,
			Margin:  60,
			Width:   1024,
		},
		DocumentStyle: styles.DocumentStyle{
			BackgroundColor: "#FFFFFF",
		},
		Identifier: identifier.String(),
	}
	a.WithDefaultStyles()
	return a
}

func (a *Article) WithDefaultStyles() {
	componentTextStyleBytes, err := ioutil.ReadFile("./configs/default_component_text_styles.json")
	if err != nil {
		panic(err)
	}
	var componentTextStyles map[string]styles.ComponentTextStyle
	if err := json.Unmarshal(componentTextStyleBytes, &componentTextStyles); err != nil {
		panic(err)
	}

	a.ComponentTextStyles = componentTextStyles

	componentLayoutBytes, err := ioutil.ReadFile("./configs/default_component_layouts.json")
	if err != nil {
		panic(err)
	}
	var componentLayouts map[string]layouts.ComponentLayout
	if err := json.Unmarshal(componentLayoutBytes, &componentLayouts); err != nil {
		panic(err)
	}
	a.ComponentLayouts = componentLayouts
}

func (article *Article) PushComponent(component components.Component) *Article {
	article.Components = append(article.Components, component)
	return article
}

func (artcle *Article) String() string {
	jsonBytes, err := json.Marshal(artcle)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
