package main

import (
	"time"
)

/**
*
 */
type PrepareQueue struct {
	Queue     []*Task
	WorkQueue chan *Task
}

/**
*
 */
func (pq *PrepareQueue) Add(task *Task) {
	// 新的任务入队
	pq.Queue = append(pq.Queue, task)
}

/**
*
 */
func (pq *PrepareQueue) Delete(task *Task) {
	queue := make([]*Task, 0)

	for _, t := range pq.Queue {
		if t.Name != task.Name {
			queue = append(queue, t)
		}
	}

	pq.Queue = queue
}

/**
*
 */
func (pq *PrepareQueue) Exists(task *Task) bool {
	for _, t := range pq.Queue {
		if t.Name == task.Name {
			return true
		}
	}
	return false
}

/**
*
 */
func (pq *PrepareQueue) Update() {
	for {
		pq.Check()
		time.Sleep(1 * time.Second)
	}
}

/**
*
 */
func (pq *PrepareQueue) Check() {
	for _, task := range pq.Queue {
		if task.Count%int(task.Interval/time.Second) == 0 {
			// 放入队列
			pq.WorkQueue <- task
		}
		task.Count++
	}
}
