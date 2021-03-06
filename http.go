package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	runewidth "github.com/mattn/go-runewidth"
)

type job struct {
	name     string
	duration time.Duration
}

func doWork(id int, t *Task) {
	t.Update()
	// fmt.Printf("worker%d: started %s, working for %f seconds\n", id, j.name, j.duration.Seconds())
	// time.Sleep(j.duration)
	// fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func requestHandler(jobs chan *Task, w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the durations.
	interval, err := time.ParseDuration(r.FormValue("interval"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set name and validate value.
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "You must specify a rss url.", http.StatusBadRequest)
		return
	}

	// Create Job and push the work onto the jobCh.
	task := &Task{Name: name, URL: url, Interval: interval}
	go func() {
		jobs <- task
	}()

	// Render success.
	w.WriteHeader(http.StatusCreated)
	return
}

func addHandler(pq *PrepareQueue, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	url := r.FormValue("url")
	exclude := r.FormValue("exclude")

	fmt.Println(url, r.FormValue("interval"))
	interval, err := time.ParseDuration(r.FormValue("interval"))
	if err != nil {
		http.Error(w, "Bad interval value: "+err.Error(), http.StatusBadRequest)
		return
	}

	task := &Task{Name: name, URL: url, Interval: interval, Exclude: exclude}
	if pq.Exists(task) {
		println("已经存在该任务")
	} else {
		pq.Add(task)
		fmt.Println("添加了任务", task.Name)
	}

	// Render success.
	w.WriteHeader(http.StatusCreated)
	return
}

func fillRight(s string, w int) string {
	return runewidth.FillRight(s, w)
}

func listHandler(pq *PrepareQueue, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s%s\n", fillRight("Name", 35), fillRight("URL", 55), fillRight("Interval", 30))
	for _, v := range pq.Queue {
		fmt.Fprintf(w, "%s%s%s\n", fillRight(v.Name, 35), fillRight(v.URL, 55), fillRight(v.Interval.String(), 30))
	}
}

func deleteHandler(pq *PrepareQueue, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	task := &Task{Name: name}
	if pq.Exists(task) {
		pq.Delete(task)
		fmt.Printf("%s 任务删除成功\n", task.Name)
	} else {
		fmt.Println("没有该任务")
	}
}

/**
*
 */
func RunHTTP(options map[string]string) {
	const (
		maxQueueSize = 100
		maxWorkers   = 5
	)
	// create job channel
	jobs := make(chan *Task, maxQueueSize)
	pq := &PrepareQueue{WorkQueue: jobs}

	// create workers
	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			for j := range jobs {
				doWork(i, j)
			}
		}(i)
	}

	go pq.Update()
	// handler for adding jobs
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(jobs, w, r)
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// requestHandler(jobs, w, r)
		addHandler(pq, w, r)
	})
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// requestHandler(jobs, w, r)
		listHandler(pq, w, r)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		// requestHandler(jobs, w, r)
		deleteHandler(pq, w, r)
	})
	log.Fatal(http.ListenAndServe(":"+options["port"], nil))
}
