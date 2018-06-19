package article

import (
	"fmt"
	"encoding/json"
	"github.com/prometheus/common/log"
	"testing"
	"github.com/sdotz/apple-news-format/pkg/components"
)

func TestMakeBasicArticleJson(t *testing.T) {
	article := Article{
		Identifier: "55910234",
		Components: []components.Component{
			components.NewBody().SetText("Hi Mom!"),
		},
	}

	j, err := json.MarshalIndent(&article, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON: ", string(j))
}
