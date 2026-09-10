package manager

import (
	"fmt"
	"uuid"

	"github.com/golang-collections/collections/queue"
	"github.com/peterouob/orchestrator/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[uuid.UUID][]task.Task
	EventDb       map[uuid.UUID][]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID][]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("manager select worker ...")
}

func (m *Manager) UpdateTask() {
	fmt.Println("manager updated task ...")
}

func (m *Manager) SendWorker() {
	fmt.Println("manager run worker ...")
}
