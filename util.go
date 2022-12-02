package main

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v3"
)

func PrettyJSON(in any) string {
	b, _ := json.MarshalIndent(&in, "", "  ")
	return string(b)
}

func PrettyYAML(in any) string {
	b, err := yaml.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
