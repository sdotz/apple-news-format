package convert

import (
	"encoding/json"
	"io/ioutil"
)

type SiteConversionConfig struct {
	SectionConversionSelectors []string `json:"sectionConversionSelectors,omitempty"`
	TitleSelector              string   `json:"titleSelector"`
}

func GetSiteConfig(conversionConfigFilePath string) (*SiteConversionConfig, error) {
	b, err := ioutil.ReadFile(conversionConfigFilePath)
	if err != nil {
		return nil, err
	}
	var config SiteConversionConfig
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
