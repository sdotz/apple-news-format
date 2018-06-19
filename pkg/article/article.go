package article

import (
	"github.com/sdotz/apple-news-format/pkg/components"
	"github.com/sdotz/apple-news-format/pkg/styles"
	"github.com/sdotz/apple-news-format/pkg/layouts"
	"github.com/sdotz/apple-news-format/pkg/properties"
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
)

type Article struct {
	Version             string                               `json:"version"`
	Identifier          string                               `json:"identifier"`
	Language            string                               `json:"language"`
	Title               string                               `json:"title"`
	Subtitle            string                               `json:"subtitle"`
	Layout              Layout                               `json:"layout"`
	AdvertisingSettings AdvertisingSettings                  `json:"advertisingSettings"`
	Components          []components.Component               `json:"components"`
	DocumentStyle       styles.DocumentStyle                 `json:"documentStyle"`
	ComponentTextStyles map[string]styles.ComponentTextStyle `json:"componentTextStyles"`
	ComponentStyles     map[string]styles.ComponentStyle     `json:"componentStyles"`
	TextStyles          map[string]styles.TextStyle          `json:"textStyles"`
	ComponentLayouts    map[string]layouts.ComponentLayout   `json:"componentLayouts"`
}

type DocumentStyle struct {
	BackgroundColor string `json:"backgroundColor"`
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
	Margin properties.Margin `json:"margin"`
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
