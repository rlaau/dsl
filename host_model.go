package dsl

type HostModel interface{}
type Host struct {
	ipAddress     string
	hostId        int
	machineStates MachineState
}
type MachineState int

const (
	Created MachineState = iota
	PanicOccured
	ErrorOccured
)

type WatcherState int

const (
	Mounted WatcherState = iota
	Unmounted
	Sleeping
	Running
)
