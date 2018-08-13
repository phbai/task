package main

import (
	"bytes"
	"time"

	"github.com/alecthomas/template"
	"github.com/phbai/task/util"
)

/**
*
 */
type Post struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	URL      string        `json:"url"`
	Interval time.Duration `json:"interval"`
	Handler  string        `json:"handler"`
	Status   string        `json:"status"`
}

/**
*
 */
type GraqhQL struct {
}

/**
*
 */
func (gq GraqhQL) MakeListGraqhQL() util.GraqhQL {
	graphqlTemplate := `{"query":"{\n  task {\n    id\n    name\n    interval\n    handler\n    url\n  }\n}","variables":null}`
	tmpl, err := template.New("graphql").Parse(graphqlTemplate)
	buf := new(bytes.Buffer)

	if err != nil {
		panic(err)
	}
	tmpl.Execute(buf, nil)

	return util.GraqhQL(buf.String())
}

/**
*
 */
func (gq GraqhQL) Add(p Post) {

}

/**
*
 */
func (gq GraqhQL) MakeAddGraqhQL(p Post) util.GraqhQL {
	graphqlTemplate := `{"query":"mutation insert_task {\n  insert_task(\n    objects: [\n      {\n        name: \"{{.Name}}\\",\n        url: \"{{.URL}}\\",\n        interval: \"{{.Interval}}\\",\n        handler: \"{{.Handler}}\\"\n      }\n    ]\n  ) {\n    returning {\n      id\n    }\n  }\n}","variables":null,"operationName":"insert_task"}`
	tmpl, err := template.New("graphql").Parse(graphqlTemplate)
	buf := new(bytes.Buffer)

	if err != nil {
		panic(err)
	}
	tmpl.Execute(buf, p)

	return util.GraqhQL(buf.String())
}

/**
*
 */
func (gq GraqhQL) MakeDeleteGraqhQL(id int) util.GraqhQL {
	graphqlTemplate := `{"query":"mutation delete_task {\n  delete_task(\n    where: {id: {_eq: {{.}}}}\n  ) {\n    affected_rows\n  }\n}","variables":null,"operationName":"delete_task"}`
	tmpl, err := template.New("graphql").Parse(graphqlTemplate)
	buf := new(bytes.Buffer)

	if err != nil {
		panic(err)
	}
	tmpl.Execute(buf, id)

	return util.GraqhQL(buf.String())

}
