package convert

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/sdotz/apple-news-format/pkg/article"
	"github.com/sdotz/apple-news-format/pkg/components"
	"golang.org/x/net/html"
)

type LinkFormatter interface {
	Format(string) string
}

type DefaultLinkFormatter struct{}

type CustomComponentHandler interface {
	Handle(*goquery.Selection) ([]components.Component, error)
	Matches(*goquery.Selection) bool
}

type Converter struct {
	LinkFormatter           LinkFormatter
	CustomComponentHandlers []CustomComponentHandler
}

func (df *DefaultLinkFormatter) Format(inHREF string) string {
	return inHREF
}

func (c *Converter) HTMLToANF(url string, htmlBytes []byte, siteConfig *SiteConversionConfig) (*article.Article, error) {
	myArticle := article.NewArticleWithDefaults()
	//Example from https://godoc.org/golang.org/x/net/html#example-Parse

	components, err := c.HTMLToANFComponents(htmlBytes, siteConfig)
	if err != nil {
		return nil, err
	}

	myArticle.Components = components
	//myArticle.Title = doc.Find(siteConfig.TitleSelector).Text()
	myArticle.Metadata.CanonicalURL = url
	myArticle.Version = article.FORMAT_VERSION_1_7
	myArticle.Language = "en"

	return &myArticle, nil
}

func (c *Converter) HTMLToANFComponents(htmlBytes []byte, siteConfig *SiteConversionConfig) ([]components.Component, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
	if err != nil {
		return nil, err
	}

	var components []components.Component
	if siteConfig != nil && len(siteConfig.SectionConversionSelectors) > 0 {
		for _, v := range siteConfig.SectionConversionSelectors {
			selection := doc.Find(v)
			components, err = c.bodyBuilderFunction(components, selection)
			if err != nil {
				return components, err
			}
		}
	} else {
		components, err = c.bodyBuilderFunction(components, doc.Slice(0, goquery.ToEnd))
		if err != nil {
			return components, err
		}
	}
	return components, nil
}

func getElementAttr(element *html.Node, attr string) *html.Attribute {
	var retVal *html.Attribute
	for _, a := range element.Attr {
		if a.Key == attr {
			retVal = &a
			break
		}
	}
	return retVal
}

func renderHTML(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)

	html.Render(w, n)
	return buf.String()
}

func (converter *Converter) bodyBuilderFunction(cs []components.Component, n *goquery.Selection) ([]components.Component, error) {
	for _, v := range converter.CustomComponentHandlers {
		if v.Matches(n) {
			if customComponent, err := v.Handle(n); err == nil {
				cs = append(cs, customComponent...)
			}
		}
	}
	switch goquery.NodeName(n) {
	case "p":
		if len(n.Text()) == 0 {
			break
		}
		var buf bytes.Buffer
		w := io.Writer(&buf)
		paragraph := components.NewBody()
		//TODO: is picking the first one here wrong?
		converter.appendTrackingParams(n.Get(0))
		//TODO: and here...
		html.Render(w, n.Get(0))
		paragraph.SetText(buf.String())
		paragraph.SetFormat(components.FormatHtml)
		paragraph.SetLayout("default-body")
		cs = append(cs, paragraph)
		break
	case "h1":
		h := components.NewHeadingWithLevel(components.HEADINGLEVEL1)
		h.SetText(renderHTML(n.Get(0)))
		h.SetFormat(components.FormatHtml)
		cs = append(cs, h)
		break
	case "h2":
		h := components.NewHeadingWithLevel(components.HEADINGLEVEL2)
		h.SetText(renderHTML(n.Get(0)))
		h.SetFormat(components.FormatHtml)
		cs = append(cs, h)
		break
	case "h3":
		h := components.NewHeadingWithLevel(components.HEADINGLEVEL3)
		h.SetText(renderHTML(n.Get(0)))
		h.SetFormat(components.FormatHtml)
		cs = append(cs, h)
		break
	case "h4":
		h := components.NewHeadingWithLevel(components.HEADINGLEVEL4)
		h.SetText(renderHTML(n.Get(0)))
		h.SetFormat(components.FormatHtml)
		cs = append(cs, h)
		break
	case "h5":
		h := components.NewHeadingWithLevel(components.HEADINGLEVEL5)
		h.SetText(renderHTML(n.Get(0)))
		h.SetFormat(components.FormatHtml)
		cs = append(cs, h)
		break
	case "h6":
		h := components.NewHeadingWithLevel(components.HEADINGLEVEL6)
		h.SetText(renderHTML(n.Get(0)))
		h.SetFormat(components.FormatHtml)
		cs = append(cs, h)
		break
	case "img":
		//The Image component has pinch to zoom disabled.
		image := components.NewImage()
		src, exists := n.Attr("src")
		if !exists {
			break
		}

		if _, err := url.ParseRequestURI(src); err != nil {
			errLog.Println(err.Error())
			break
		}

		image.URL = src

		if altText, exists := n.Attr("alt"); exists {
			//don't add it if it is only whitespace
			if !isWhitespace(altText) {
				image.AccessibilityCaption = altText
			}
		}
		cs = append(cs, image)
		break
	case "video":
		video := components.NewVideo()
		src, exists := n.Attr("src")
		if !exists {
			break
		}

		if _, err := url.ParseRequestURI(src); err != nil {
			errLog.Println(err.Error())
			break
		}

		video.URL = src

		cs = append(cs, video)
		break
	case "figcaption":
		var buf bytes.Buffer
		w := io.Writer(&buf)
		caption := components.NewCaption()
		caption.SetFormat(components.FormatHtml)
		html.Render(w, n.Get(0))
		if !isWhitespace(buf.String()) {
			caption.Text = buf.String()
			cs = append(cs, caption)
		}
		break
	case "ul":
		list := components.NewBody()
		//TODO: get spacing between list items
		list.SetLayout("default-body")
		list.SetFormat(components.FormatHtml)
		var buf bytes.Buffer
		w := io.Writer(&buf)
		// render the root node (ul) and all <li> will render with it
		html.Render(w, n.Get(0))
		list.SetText(buf.String())
		cs = append(cs, list)
		break
	default:
	}

	for c := n.Children().First(); c.Length() > 0; c = c.Next() {
		var childComponents []components.Component
		childComponents, err := converter.bodyBuilderFunction(childComponents, c)
		if err != nil {
			return cs, err
		}
		cs = append(cs, childComponents...)
	}
	return cs, nil
}

func (converter *Converter) appendTrackingParams(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for k, a := range n.Attr {
			if a.Key == "href" {
				n.Attr[k].Val = converter.LinkFormatter.Format(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		converter.appendTrackingParams(c)
	}
}

func isWhitespace(testString string) bool {
	return regexp.MustCompile(`^\s+$`).Match([]byte(testString))
}

func (c *Converter) ConvertUrlToAnf(url string, siteConfig *SiteConversionConfig) (*article.Article, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("Got non 200 response %d", resp.StatusCode)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return c.HTMLToANF(url, bodyBytes, siteConfig)
}
