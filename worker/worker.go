package worker

import (
	"fmt"
	"uuid"

	"github.com/golang-collections/collections/queue"
	"github.com/peterouob/orchestrator/task"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) RunTask() {
	fmt.Println("run task")
}

func (w *Worker) CollectionState() {
	fmt.Println("collection state")
}

func (w *Worker) StartTask() {
	fmt.Println("start task ...")
}

func (w *Worker) StopTask() {
	fmt.Println("stop task ...")
}
