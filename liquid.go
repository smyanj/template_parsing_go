package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/osteele/liquid"
)

func main() {
	url := ""
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get file from web guide: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read file content: %v\n", err)
	}

	template := string(body)

	engine := liquid.NewEngine()

	// Define bindings
	bindings := map[string]any{
		"page": map[string]string{
			"title": "Introduction",
		},
	}

	out, err := engine.ParseAndRenderString(template, bindings)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(out)
}
