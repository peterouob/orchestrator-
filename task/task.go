package task

import (
	"time"
	"uuid"

	"github.com/docker/go-connections/nat"
)

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type TaskEvent struct {
	Id        uuid.UUID
	State     State
	TimeStamp time.Time
	Task      Task
}

type Task struct {
	Id            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.Port
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}
