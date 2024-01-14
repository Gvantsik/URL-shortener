package main

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return

		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var data []PathURL
	err := yaml.Unmarshal(yml, &data)

	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range data {
		pathsToUrls[pu.Path] = pu.URL
	}

	return mapHandler(pathsToUrls, fallback), nil
}
