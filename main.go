package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/linkedin-gvantsa": "https://www.linkedin.com/in/gvantsa-tsotskolauri-97aa9b18b/",
		"/github-gvantsa":   "https://github.com/Gvantsik?tab=repositories",
	}
	mapHandler := mapHandler(pathsToUrls, mux)

	yaml := `
  - path: /linkedin-gvantsa
    url: https://www.linkedin.com/in/gvantsa-tsotskolauri-97aa9b18b/
  - path: /github-gvantsa
    url: https://github.com/Gvantsik?tab=repositories
  `
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
