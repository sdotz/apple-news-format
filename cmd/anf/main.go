package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sdotz/apple-news-format/pkg/convert"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	convertCmd     = kingpin.Command("convert", "Convert content to apple news format")
	url            = convertCmd.Flag("url", "A URL to download and convert to apple news format").String()
	siteConfigPath = convertCmd.Flag("siteConfig", "Path of site config json file").String()
)

func main() {
	cmd := kingpin.Parse()

	switch cmd {
	case "convert":
		siteConfig, err := convert.GetSiteConfig(*siteConfigPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if siteConfig == nil {
			siteConfig = &convert.SiteConversionConfig{
				SectionConversionSelectors: []string{"body"},
			}
		}

		article, err := convert.ConvertUrlToAnf(*url, siteConfig)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b, err := json.Marshal(article)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(b))
	}
}
