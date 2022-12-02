package main

func init() {
	puzzles = append(puzzles, &Day3Puzzle{})
}

type Day3Puzzle struct{}

func (day *Day3Puzzle) Name() string {
	return "day_3"
}

func (day *Day3Puzzle) Solution() (*Result, error) {
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, 3)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
