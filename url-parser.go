package parser

import (
	"strings"

	"github.com/pkg/errors"
)

// GetURLVars returns the variables from the given url and url schema
func GetURLVars(schema string, url string) (map[string]string, error) {
	schemaParts := strings.Split(schema, "/")
	urlParts := strings.Split(schema, "/")

	if len(schemaParts) != len(urlParts) {
		return nil, errors.New("URL does not match schema")
	}

	var vars map[string]string

	for i, part := range schemaParts {
		if !strings.HasPrefix(part, "{") || !strings.HasSuffix(part, "}") {
			continue
		}

		varname := strings.TrimPrefix(part, "{")
		varname = strings.TrimSuffix(part, "}")

		vars[varname] = urlParts[i]
	}

	return vars, nil
}
