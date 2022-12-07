package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day7Puzzle{})
}

type Day7Puzzle struct{}

func (day *Day7Puzzle) Name() string {
	return "day_7"
}

func (day *Day7Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, 7)
	if err != nil {
		return nil, err
	}
	fs := NewObject("/", Directory)
	fs.Populate(data[1:])

	var sizes = make(map[string]int)
	fs.DirSizes([]string{}, sizes)

	var (
		puzzleOne   = 0
		puzzleTwo   []int
		spaceNeeded = 30000000 - (70000000 - sizes["/"])
	)
	fmt.Printf("space needed to be freed: %d\n", spaceNeeded)
	for _, size := range sizes {
		if size < 100000 {
			puzzleOne += size
		}
		if size >= spaceNeeded {
			puzzleTwo = append(puzzleTwo, size)
		}
	}
	sort.Ints(puzzleTwo)

	return &Result{
		First:    puzzleOne,
		Second:   puzzleTwo[0],
		Duration: time.Now().Sub(begin),
	}, nil
}

type Type string

const (
	File      Type = "file"
	Directory Type = "dir"
)

type Object struct {
	Name     string
	Type     Type
	Children map[string]*Object
	size     int
}

func (o *Object) Populate(data []string) []string {
	var (
		dirRegex  = regexp.MustCompile(`dir (\w+)`)
		fileRegex = regexp.MustCompile(`(\d+) ([\w\.]+)`)
		cdRegex   = regexp.MustCompile(`\$ cd (.*)`)
	)
	for i, line := range data {
		if line == `$ ls` {
			continue
		}

		if dirRegex.MatchString(line) {
			match := dirRegex.FindStringSubmatch(line)
			objName := match[1]
			o.Children[objName] = NewObject(objName, Directory)
			continue
		}
		if fileRegex.MatchString(line) {
			match := fileRegex.FindStringSubmatch(line)
			objName := match[2]
			o.Children[objName] = NewObject(objName, File)
			o.Children[objName].size, _ = strconv.Atoi(match[1])
			continue
		}
		if cdRegex.MatchString(line) {
			match := cdRegex.FindStringSubmatch(line)
			switch match[1] {
			case "..":
				return data[i+1:]
			default:
				return o.Populate(o.Children[match[1]].Populate(data[i+1:]))
			}
		}
		log.Fatalf("I dunno how we got here: %v", line)
	}
	return nil
}

func NewObject(name string, t Type) *Object {
	return &Object{
		Name:     name,
		Type:     t,
		Children: make(map[string]*Object),
	}
}

func (o *Object) Size() int {
	if o.Type == File {
		return o.size
	}
	var sum = o.size
	for _, c := range o.Children {
		sum += c.Size()
	}
	return sum
}

func (o *Object) DirSizes(parentPath []string, sizes map[string]int) {
	parentPath = append(parentPath, o.Name)
	path := strings.Join(parentPath, "/")
	sizes[path] = o.Size()
	for _, c := range o.Children {
		if c.Type == File {
			continue
		}
		c.DirSizes(parentPath, sizes)
	}
}
