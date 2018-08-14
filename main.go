package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	task = kingpin.New("task", "A rss task runner.")

	daemon     = task.Command("daemon", "start the daemon")
	daemonPort = daemon.Flag("port", "daemon port. dafault port is 8080").Default("8080").String()

	list    = task.Command("list", "list the running tasks")
	listAll = list.Flag("all", "list all the tasks.").Default("false").Bool()

	add         = task.Command("add", "add a task to the task queue")
	addName     = add.Flag("name", "task name").Required().String()
	addURL      = add.Flag("url", "task rss url").Required().String()
	addInterval = add.Flag("interval", "task interval").Default("300s").Duration()
	// addHandler  = add.Flag("handler", "task handler").Default("handler.go").String()

	delete   = task.Command("delete", "delete a task by id")
	deleteID = delete.Flag("id", "task id.").Required().Int()
)

func main() {

	switch kingpin.MustParse(task.Parse(os.Args[1:])) {

	case daemon.FullCommand():
		options := make(map[string]string)
		options["port"] = *daemonPort
		fmt.Printf("listen on the http://localhost:%s\n", options["port"])
		RunHTTP(options)

	case list.FullCommand():
		// graqhql := gq.MakeListGraqhQL()
		// util.GraphQLPost(graqhql)

	case add.FullCommand():
		// p := Post{Name: *addName, URL: *addURL, Interval: *addInterval, Handler: *addHandler}
		// task := &Task{Name: *addName, URL: *addURL, Interval: *addInterval, Count: 0}
		// graqhql := gq.MakeAddGraqhQL(p)
		// util.GraphQLPost(graqhql)
		// fmt.Println(*addName, *addURL, *addInterval, *addHandler)

	case delete.FullCommand():
		// graqhql := gq.MakeDeleteGraqhQL(*deleteID)
		// util.GraphQLPost(graqhql)
	}
}
