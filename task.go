package main

import (
	"fmt"
	"time"

	"github.com/SlyMarbo/rss"
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
func (t *Task) CheckUpdate() {
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
	for _, value := range t.feed.Items {
		if !value.Read {
			fmt.Printf("%s更新了 %s 地址: %s\n", t.Name, value.Title, value.Link)
		}
		value.Read = true
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
