package main

import (
	"io"
	"net/http"

	"github.com/skx/gobasic/builtin"
	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

func main() {
	code := `
	REM custom function
	res$ = HTTPGET ("https://httpbin.org/anything")
	PRINT res$
	`
	tok := tokenizer.New(code)

	interp, err := eval.New(tok)
	if err != nil {
		panic(err)
	}

	httpGetFunc := func(env builtin.Environment, args []object.Object) object.Object {
		url := args[0].(*object.StringObject).Value
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return &object.StringObject{Value: string(body)}
	}

	interp.RegisterBuiltin("HTTPGET", 1, httpGetFunc)
	err = interp.Run()
	if err != nil {
		panic(err)
	}
}
