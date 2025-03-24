package parser

import (
	"flag"
	"fmt"
)

func ParseDiff() (string, error) {
	diff := flag.String("diff", "", "diff to parse")
	flag.Parse()

	if *diff == "" {
		return "", fmt.Errorf("diff is required")
	}

	return *diff, nil
}
