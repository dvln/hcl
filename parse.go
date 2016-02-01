package hcl

import (
	"fmt"

	"github.com/dvln/hcl/hcl/ast"
	hclParser "github.com/dvln/hcl/hcl/parser"
	jsonParser "github.com/dvln/hcl/json/parser"
)

// Parse parses the given input and returns the root object.
//
// The input format can be either HCL or JSON.
func Parse(input string) (*ast.File, error) {
	switch lexMode(input) {
	case lexModeHcl:
		return hclParser.Parse([]byte(input))
	case lexModeJson:
		return jsonParser.Parse([]byte(input))
	}

	return nil, fmt.Errorf("unknown config format")
}
