package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type input struct {
	condition string
	data      interface{}
}

func main() {
	inps := []input{
		{`.Foo == "foo"`, map[string]string{"Foo": "foo"}},
		{`eq .Foo "foo"`, map[string]string{"Foo": "foo"}},
		{`eq .Foo "bar"`, map[string]string{"Foo": "foo"}},
		{`not (eq .Foo "foo")`, map[string]string{"Foo": "foo"}},
		{`not (eq .Foo "bar")`, map[string]string{"Foo": "foo"}},
		{`eq .Foo .Bar`, map[string]string{"Foo": "aaa", "Bar": "aaa"}},
		{`eq .Foo .Bar`, map[string]string{"Foo": "aaa", "Bar": "bbb"}},
		{`lt (len .List) 3`, map[string][]string{"List": {"a", "b"}}},
		{`lt (len .List) 3`, map[string][]string{"List": {"a", "b", "c"}}},
		{`lt (len .List) 3`, map[string][]string{"List": {"a", "b", "c", "d"}}},
	}

	for _, inp := range inps {
		fmt.Println("data:     ", inp.data)
		fmt.Println("condition:", inp.condition)
		result, err := TemplateConditionEval(inp.condition, inp.data)
		if err != nil {
			fmt.Println("err:      ", err)
		} else {
			fmt.Println("result:   ", result)
		}
		fmt.Println()
	}
}

func TemplateConditionEval(condition string, data interface{}) (bool, error) {
	tmpl, err := template.New("condition-eval").Parse(`{{if ` + condition + `}}true{{end}}`)
	if err != nil {
		return false, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return false, err
	}

	return buf.String() == "true", nil
}
