package convert

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sdotz/apple-news-format/pkg/components"
	"net/http"
	"io/ioutil"
	"github.com/sdotz/apple-news-format/pkg/article"
	"github.com/pkg/errors"
	"bytes"
	"golang.org/x/net/html"
	url2 "net/url"
)

func ConvertHtmlToAnf(url string, htmlBytes []byte, siteConfig *SiteConversionConfig) (*article.Article, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
	if err != nil {
		return nil, err
	}

	myArticle := article.NewArticleWithDefaults()
	//Example from https://godoc.org/golang.org/x/net/html#example-Parse
	for _, v := range siteConfig.SectionConversionSelectors {
		selection := doc.Find(v)
		bodyBuilderFunction(&myArticle, selection.Get(0))
	}

	myArticle.Title = doc.Find(siteConfig.TitleSelector).Text()
	myArticle.Metadata.CanonicalURL = url
	myArticle.Version = article.FORMAT_VERSION_1_7
	myArticle.Language = "en"

	return &myArticle, nil
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

func bodyBuilderFunction(article *article.Article, n *html.Node) error {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "p":
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				paragraph := components.NewBody()
				paragraph.SetText(n.FirstChild.Data)
				article.PushComponent(paragraph)
			}
			break
		case "h1":
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				h := components.NewHeadingWithLevel(components.HEADINGLEVEL1)
				h.SetText(n.FirstChild.Data)
				article.PushComponent(h)
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

			article.PushComponent(image)
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

			article.PushComponent(video)
			break
		default:
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		bodyBuilderFunction(article, c)
	}
	return nil
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
