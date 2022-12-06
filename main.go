package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/shurcooL/markdownfmt/markdown"
)

type Puzzle interface {
	Name() string
	Solution() (*Result, error)
}

type Result struct {
	First    any
	Second   any
	Duration time.Duration
}

func (r *Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		First    any    `json:"first"`
		Second   any    `json:"second"`
		Duration string `json:"duration"`
	}{
		First:    r.First,
		Second:   r.Second,
		Duration: r.Duration.String(),
	})
}

var puzzles []Puzzle

func main() {
	var totalDuration int64
	results := make(map[string]*Result)
	for _, p := range puzzles {
		solution, err := p.Solution()
		if err != nil {
			log.Fatalf("failed to solve puzzle %s: %v", p.Name(), err)
			continue
		}
		results[p.Name()] = solution
		totalDuration += solution.Duration.Nanoseconds()
	}
	tmpl, err := template.ParseFiles("README.md.tpl")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, map[string]any{
		"results":       results,
		"totalDuration": time.Duration(totalDuration),
	}); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("README.md")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	b, err := markdown.Process(f.Name(), buf.Bytes(), &markdown.Options{Terminal: false})
	if err != nil {
		log.Fatal(err)
	}

	f.Write(b)
}
