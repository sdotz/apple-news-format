package article

import (
	"fmt"
	"encoding/json"
	"github.com/prometheus/common/log"
	"testing"
	"github.com/sdotz/apple-news-format/pkg/components"
	"github.com/sdotz/apple-news-format/pkg/styles"
)

func TestMakeBasicArticleJson(t *testing.T) {
	article := Article{
		Identifier: "55910234",
		Components: []components.Component{
			components.NewTitle().SetText("This is the title!").SetFormat(components.FormatMarkdown),
			components.NewBody().SetText("Welcome to my story").SetStyle(&styles.ComponentStyle{
				BackgroundColor: styles.Color("#ffffff"),
			},),
		},
	}

	j, err := json.MarshalIndent(&article, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON: ", string(j))
}
