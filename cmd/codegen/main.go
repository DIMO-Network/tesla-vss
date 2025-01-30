package main

import (
	"log"

	"github.com/DIMO-Network/tesla-vss/pkg/codegen"
)

const (
	defaultPackage        = "convert"
	defaultRulesPath      = "pkg/schema/schema.yaml"
	defaultOuterFuncsPath = "pkg/convert/outer_convert_funcs_gen.go"
	defaultInnerFuncsPath = "pkg/convert/inner_convert_funcs_gen.go"
)

func main() {
	err := codegen.Generate(defaultPackage, defaultRulesPath, defaultOuterFuncsPath, defaultInnerFuncsPath)
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}
}
