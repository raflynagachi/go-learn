package helper

import (
	"encoding/json"
	"urlshort/model"

	"github.com/go-yaml/yaml"
)

func ParseYaml(yml []byte, pathUrls *[]model.PathToUrl) {
	err := yaml.Unmarshal(yml, pathUrls)
	if err != nil {
		panic(err)
	}
}

func ParseJson(jsonBytes []byte, pathUrls *[]model.PathToUrl) {
	err := json.Unmarshal(jsonBytes, pathUrls)
	if err != nil {
		panic(err)
	}
}

func BuildMap(pathUrls []model.PathToUrl) map[string]string {
	var pathToUrls = make(map[string]string)
	for _, pathUrl := range pathUrls {
		pathToUrls[pathUrl.Path] = pathUrl.URL
	}
	return pathToUrls
}
