package node

type Node struct {
	Name            string
	Ip              string
	Role            string
	Cores           int
	Memory          int
	Disk            int
	MemoryAllocated int
	DiskAllocated   int
	TaskCount       int
}
