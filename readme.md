## rss runner

[![Status in progress](https://img.shields.io/badge/Status-in%20progress-orange.svg)](#)

```go
go get -u github.com/phbai/task

task daemon

/* 
* name: task name
* url: rss url
* interval: task interval
*/
curl localhost:8080/add -d name=v2ex订阅 -d url="https://www.v2ex.com/index.xml" -d interval=15s // add a task
curl localhost:8080/delete -d name=v2ex订阅 // delete the specfic task
curl localhost:8080/list // list all the task
```

### TODO:
- [ ] add webui
- [ ] support mongodb updated item will be inserted into mongodb automatically
- [ ] add docker、docker-compose support
- [ ] add task cli：task delete、task add、task list
