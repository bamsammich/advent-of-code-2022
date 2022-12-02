package main

import (
	"encoding/json"
)

func PrettyJSON(in any) string {
	b, _ := json.MarshalIndent(&in, "", "  ")
	return string(b)
}
