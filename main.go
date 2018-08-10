package main

import (
	"os"

	"github.com/phbai/task/util"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	task = kingpin.New("task", "A rss task runner.")

	list    = task.Command("list", "list the running tasks")
	listAll = list.Flag("all", "list all the tasks.").Default("false").Bool()

	add         = task.Command("add", "add a task to the task queue")
	addName     = add.Flag("name", "task name").Required().String()
	addURL      = add.Flag("url", "task rss url").Required().String()
	addInterval = add.Flag("interval", "task interval").Default("300s").Duration()
	addHandler  = add.Flag("handler", "task handler").Default("handler.go").String()

	delete   = task.Command("delete", "delete a task by id")
	deleteID = delete.Flag("id", "task id.").Required().Int()
)

func main() {
	switch kingpin.MustParse(task.Parse(os.Args[1:])) {

	case list.FullCommand():
		t := Task{}
		graqhql := t.MakeListGraqhQL()
		util.GraphQLPost(graqhql)

	case add.FullCommand():
		t := Task{}
		p := Post{Name: *addName, URL: *addURL, Interval: *addInterval, Handler: *addHandler}
		graqhql := t.MakeAddGraqhQL(p)
		util.GraphQLPost(graqhql)
		// fmt.Println(*addName, *addURL, *addInterval, *addHandler)

	case delete.FullCommand():
		t := Task{}
		graqhql := t.MakeDeleteGraqhQL(*deleteID)
		util.GraphQLPost(graqhql)
	}
}

// package main

// import (
// 	"os"
// )

// func main() {
// 	task := Task{}
// 	args := os.Args[1:]
// 	switch args[0] {
// 	case "list":
// 		task.List()
// 	case "delete":
// 		task.Delete()
// 	}
// }
