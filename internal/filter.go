package internal

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

var celEnv *cel.Env

func CreateCELProgram(filter string) (cel.Program, error) {
	ast, issues := celEnv.Compile(filter)
	if issues != nil && issues.Err() != nil {
		return nil, issues.Err()
	}
	return celEnv.Program(ast)
}

func init() {
	var err error
	headerType := decls.NewMapType(decls.String, decls.String)
	celEnv, err = cel.NewEnv(
		cel.Declarations(
			decls.NewVar("host", decls.String),
			decls.NewVar("path", decls.String),
			decls.NewVar("method", decls.String),
			decls.NewVar("source_address", decls.String),
			decls.NewVar("header", headerType),
		),
	)
	if err != nil {
		panic(err)
	}
}
