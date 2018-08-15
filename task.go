package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/SlyMarbo/rss"
	"github.com/zyxar/argo/rpc"
)

/**
*
 */
type Task struct {
	Name     string
	URL      string
	Interval time.Duration
	LastName string
	Handler  func()
	Count    int
	Exclude  string
	feed     *rss.Feed
}

/**
* 工作流程：
* 从url获取feed  判断是否有feed更新
 */
func (t *Task) Start() {
	for {
		t.CheckRSS()
		time.Sleep(t.Interval)
	}
}

/**
*
 */
func (t *Task) Update() {
	t.CheckRSS()
}

/**
*
 */
func (t *Task) CheckRSS() {
	if t.feed == nil {
		t.feed, _ = rss.Fetch(t.URL)
	} else {
		t.feed.Update()
	}
	for _, item := range t.feed.Items {
		if !item.Read && !strings.Contains(item.Title, t.Exclude) {
			fmt.Printf("%s更新了 %s 地址: %s\n", t.Name, item.Title, item.Link)
			t.HandleItem(item)
		}
		item.Read = true
	}
}

/**
*
 */
func (t *Task) HandleItem(item *rss.Item) {
	for _, v := range item.Enclosures {
		p, err := rpc.New("http://localhost:6800/jsonrpc", "69378fb6a75b5d488073")
		if err != nil {
			println(err)
		}
		p.AddURI(v.URL)
	}
}

// func (t Task) GetPosts(url string) {
// 	feed, err := rss.Fetch(url)
// 	if err != nil {
// 		// handle error.
// 	}

// 	for _, v := range feed.Items {
// 		v.Read = true
// 	}

// 	time.Sleep(time.Second * 10)
// 	feed.Update()

// 	fmt.Println(feed)
// 	// feed.Update()
// 	// fmt.Println(feed)
// }
