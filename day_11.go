package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day11Puzzle{})
}

type Day11Puzzle struct{}

func (day *Day11Puzzle) Number() int {
	return 11
}

func (day *Day11Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day11Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	var (
		monkeyData = strings.Split(strings.Join(data, "\n"), "\n\n")
		monkeysOne = make([]*Monkey, len(monkeyData))
		// monkeysTwo = make([]*Monkey, len(monkeyData))
	)
	for _, mData := range monkeyData {
		one := MonkeyFromString(mData)
		// two := MonkeyFromString(mData)
		monkeysOne[one.ID] = &one
		// monkeysTwo[two.ID] = &two
	}

	for i := 0; i < 10000; i++ {
		if i < 20 {
			for i := range monkeysOne {
				monkeysOne[i].Throw(monkeysOne, 3)
			}
		}
		// for i := range monkeysTwo {
		// 	monkeysTwo[i].Throw(monkeysTwo, 1)
		// }
	}

	var (
		puzzleOne []int
		// puzzleTwo []int
	)
	for i := range monkeysOne {
		puzzleOne = append(puzzleOne, monkeysOne[i].InspectCount)
		// puzzleTwo = append(puzzleTwo, monkeysTwo[i].InspectCount)
	}

	sort.Ints(puzzleOne)
	return &Result{
		First:    product(puzzleOne[len(puzzleOne)-2:]),
		Second:   nil,
		Duration: time.Now().Sub(begin),
	}, nil
}

type Monkey struct {
	ID               int
	Items            []int
	ObserveFunc      func(float64) float64 `json:"-"`
	ThrowDivisor     int
	ThrowTrueTarget  int
	ThrowFalseTarget int
	InspectCount     int
}

func (m *Monkey) String() string {
	return PrettyJSON(*m)
}

func (m *Monkey) Catch(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) Throw(monkeys []*Monkey, boredomDivisor float64) {
	for len(m.Items) > 0 {
		for i := 0; i < len(m.Items); i++ {
			m.InspectCount++
			newValue := math.Floor(m.ObserveFunc(float64(m.Items[i])) / boredomDivisor)
			m.Items[i] = int(newValue)
			if m.Items[i]%m.ThrowDivisor == 0 {
				monkeys[m.ThrowTrueTarget].Catch(m.Items[i])
			} else {
				monkeys[m.ThrowFalseTarget].Catch(m.Items[i])
			}
			m.Items = append(m.Items[:i], m.Items[i+1:]...)
		}
	}
}

func Add(a, b int) int {
	return a + b
}

func MonkeyFromString(data string) Monkey {
	m := Monkey{
		Items:        []int{},
		InspectCount: 0,
	}
	for i, field := range strings.Split(data, "\n") {
		parts := strings.Split(field, " ")
		switch i {
		case 0:
			m.ID, _ = strconv.Atoi(regexp.MustCompile(`Monkey (\d+):`).FindStringSubmatch(field)[1])
		case 1:
			for _, itemStr := range strings.Split(strings.Split(field, ": ")[1], ", ") {
				item, _ := strconv.Atoi(itemStr)
				m.Items = append(m.Items, item)
			}
		case 2:
			algParts := strings.Split(field, " ")[5:]
			if algParts[2] == "old" {
				m.ObserveFunc = func(i float64) float64 {
					return i * i
				}
				continue
			}
			factor, _ := strconv.Atoi(algParts[2])
			if algParts[1] == "+" {
				m.ObserveFunc = func(i float64) float64 {
					return i + float64(factor)
				}
			}
			if algParts[1] == "*" {
				m.ObserveFunc = func(i float64) float64 {
					return i * float64(factor)
				}
			}
		case 3:
			m.ThrowDivisor, _ = strconv.Atoi(parts[len(parts)-1])
		case 4:
			m.ThrowTrueTarget, _ = strconv.Atoi(parts[len(parts)-1])
		case 5:
			m.ThrowFalseTarget, _ = strconv.Atoi(parts[len(parts)-1])
		}
	}
	return m
}
