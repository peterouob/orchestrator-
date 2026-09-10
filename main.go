package main

import (
	"fmt"
	"time"
	"uuid"

	"github.com/golang-collections/collections/queue"
	"github.com/peterouob/orchestrator/manager"
	"github.com/peterouob/orchestrator/node"
	"github.com/peterouob/orchestrator/task"
	"github.com/peterouob/orchestrator/worker"
)

func main() {
	t := task.Task{
		Id:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "image-1",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		Id:        uuid.New(),
		State:     task.Pending,
		TimeStamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task : %v\n", t)
	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Name:  "worker-1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	fmt.Printf("worker: %v\n", w)
	w.CollectionState()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[uuid.UUID][]task.Task),
		EventDb: make(map[uuid.UUID][]task.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Printf("manager : %v\n", m)
	m.SelectWorker()
	m.UpdateTask()
	m.SendWorker()

	n := node.Node{
		Name:   "node-1",
		Ip:     "192.168.0.1",
		Cores:  4,
		Memory: 1024,
		Disk:   256,
		Role:   "worker",
	}

	fmt.Printf("%v\n", n)
}
