package host

import "github.com/rlaaudgjs5638/dsk/watcher"

// HostModel은 Watcher를 관리하고 실행하는 엔진에 대한 명세임
type HostModel interface {
	MountWatcher(w watcher.Watcher, pw watcher.StoragePassword) error
	UnmountWathcer(wid watcher.WatcherId, pw watcher.StoragePassword) error
	StopWatcher(wid watcher.WatcherId, pw watcher.StoragePassword) error
	ResumeWatcher(wid watcher.WatcherId, pw watcher.StoragePassword) error
	CommandToWatcer(wid watcher.WatcherId, cmd string, pw watcher.StoragePassword) error

	forceXXX(wid watcher.WatcherId, cmd string) error // 비밀번호 없이 강제로 실행시키는 옵션

	checkIfSameUserIdExist(uid watcher.StorageID) bool //Host내에선 userId중복 불가임
}

// Host는 Watcher를 관리하고 실행하는 엔진
type Host struct {
	id            HostId
	ipAddress     string
	machineStates MachineState
}
type HostId int

type MachineState int

const (
	Created MachineState = iota
	Running
	PanicOccured
)
