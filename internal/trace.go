package internal

import (
	"encoding/json"
	"os"
)

func LoadTraceFromPath(path string) (Trace, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Trace{}, err
	}

	var trace Trace
	err = json.Unmarshal(b, &trace)
	if err != nil {
		return Trace{}, err
	}
	return trace, nil
}
