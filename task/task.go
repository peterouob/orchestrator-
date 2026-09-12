package task

import (
	"context"
	"io"
	"log"
	"math"
	"os"
	"time"
	"uuid"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/client"
	"github.com/moby/moby/pkg/stdcopy"
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

type Docker struct {
	Client *client.Client
	Config Config
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(ctx, d.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	io.Copy(os.Stdout, reader)

	rp := container.RestartPolicy{
		Name: container.RestartPolicyDisabled,
	}

	r := container.Resources{
		Memory:   d.Config.Memory,
		NanoCPUs: int64(d.Config.Cpu * math.Pow(10, 9)),
	}

	cc := container.Config{
		Image:        d.Config.Image,
		Tty:          false,
		Env:          d.Config.Env,
		ExposedPorts: d.Config.ExposedPorts,
	}

	hc := container.HostConfig{
		RestartPolicy:   rp,
		Resources:       r,
		PublishAllPorts: true,
	}

	resp, err := d.Client.ContainerCreate(ctx, &cc, &hc, nil, nil, d.Config.Name)
	if err != nil {
		log.Printf("Error crating container using unage %s : %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}

	if err = d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Printf("Error container start %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	out, err := d.Client.ContainerLogs(ctx, resp.ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		log.Printf("Error getting logs for container %s: %v \n", resp.ID, err)
		return DockerResult{Error: err}
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return DockerResult{
		ContainerId: resp.ID,
		Action:      "start",
		Result:      "success",
	}

}

func (d *Docker) Stop(id string) DockerResult {
	log.Printf("Stop the container for id: %s", id)
	ctx := context.Background()
	if err := d.Client.ContainerStop(ctx, id, container.StopOptions{}); err != nil {
		log.Printf("Error container stop %s:%v\n", id, err)
		return DockerResult{Error: err}
	}

	if err := d.Client.ContainerRemove(ctx, id, container.RemoveOptions{}); err != nil {
		log.Printf("Error container remove %s:%v\n", id, err)
		return DockerResult{Error: err}
	}

	return DockerResult{Action: "stop", ContainerId: id}
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

type Config struct {
	Name         string
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	ExposedPorts nat.PortSet
	Cmd          []string
	Image        string
	Cpu          float64
	Memory       int64
	Disk         int64
	Env          []string
	RestartPoliy string
}
