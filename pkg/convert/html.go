package convert

import (
	"bytes"
	"io/ioutil"
	"net/http"
	url2 "net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/sdotz/apple-news-format/pkg/article"
	"github.com/sdotz/apple-news-format/pkg/components"
	"golang.org/x/net/html"
)

func ConvertHtmlToAnf(url string, htmlBytes []byte, siteConfig *SiteConversionConfig) (*article.Article, error) {

	myArticle := article.NewArticleWithDefaults()
	//Example from https://godoc.org/golang.org/x/net/html#example-Parse

	components, err := HTMLToANFComponents(htmlBytes, siteConfig)
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

func HTMLToANFComponents(htmlBytes []byte, siteConfig *SiteConversionConfig) ([]components.Component, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
	if err != nil {
		return nil, err
	}

	var components []components.Component
	if siteConfig != nil && len(siteConfig.SectionConversionSelectors) > 0 {
		for _, v := range siteConfig.SectionConversionSelectors {
			selection := doc.Find(v)
			components, err = bodyBuilderFunction(components, selection.Get(0))
			if err != nil {
				return components, err
			}
		}
	} else {
		components, err = bodyBuilderFunction(components, doc.Get(0))
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

func bodyBuilderFunction(cs []components.Component, n *html.Node) ([]components.Component, error) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "p":
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				paragraph := components.NewBody()
				paragraph.SetText(n.FirstChild.Data)
				paragraph.SetFormat(components.FormatHtml)
				paragraph.SetLayout("default-body")
				cs = append(cs, paragraph)
			}
			break
		case "h1":
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				h := components.NewHeadingWithLevel(components.HEADINGLEVEL1)
				h.SetText(n.FirstChild.Data)
				cs = append(cs, h)
			}
			break
		case "img":
			image := components.NewImage()
			url := getElementAttr(n, "src")
			if _, err := url2.ParseRequestURI(url.Val); err != nil {
				errLog.Println(err.Error())
				break
			}

			if url != nil {
				image.URL = url.Val
			}

			altText := getElementAttr(n, "alt")
			if altText != nil {
				image.AccessibilityCaption = altText.Val
			}

			cs = append(cs, image)
			break
		case "video":
			video := components.NewVideo()
			url := getElementAttr(n, "src")
			if _, err := url2.ParseRequestURI(url.Val); err != nil {
				errLog.Println(err.Error())
				break
			}

			if url != nil {
				video.URL = url.Val
			}

			cs = append(cs, video)
			break
		default:
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		var childComponents []components.Component
		childComponents, err := bodyBuilderFunction(childComponents, c)
		if err != nil {
			return cs, err
		}
		cs = append(cs, childComponents...)
	}
	return cs, nil
}

func ConvertUrlToAnf(url string, siteConfig *SiteConversionConfig) (*article.Article, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("Got non 200 response %d", resp.StatusCode)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return ConvertHtmlToAnf(url, bodyBytes, siteConfig)
}
