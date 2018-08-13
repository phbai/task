## rss runner

[![Status in progress](https://img.shields.io/badge/Status-in%20progress-orange.svg)](#)

```go
go get -u github.com/phbai/task

task daemon

/* 
* name: 任务名
* url: rss订阅地址
* interval: rss抓取时间间隔
*/
curl localhost:8080/add -d name=v2ex订阅 -d url="https://www.v2ex.com/index.xml" -d interval=15s // 新增rss订阅任务
curl localhost:8080/delete -d name=v2ex订阅 // 删除rss订阅任务
```