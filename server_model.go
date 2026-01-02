package dsl
type ServerModel interface{}
type Server struct {
	hostsInfo map[HostId]HostInfo
}

var _ ServerModel = (*Server)(nil)

func (s *Server) DispatchWatcher(w WatcherModel) {
	panic("not implemented")
}

type HostInfo struct {
	hostId          HostId
	ipAddress       IpAddress
	holdingWatchers []WatcherId
	holdingUsers    []UserId
	holdingPolicies []PolicyHash

	performance Performance
}
type HostId int
type IpAddress string
type WatcherId int
type UserId int
type PolicyHash int

type Performance int

const (
	High Performance = iota
	Middle
	Low
)
