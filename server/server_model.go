package server

import (
	"github.com/rlaaudgjs5638/dsk/host"
	"github.com/rlaaudgjs5638/dsk/watcher"
)

// ServerModel은 서버가 해야 할 기본적인 일들을 정의함
type ServerModel interface {
	//ReeiveCodeAndAccount는 사용자로부터 계정 정보와 코드를 받음
	ReceiveCodeAndAccount(acc Account, code watcher.WatcherLanguage)
	// dispatchWatcher는 Watcher를 Host에 마운트시킴
	DispatchWatcher(w watcher.WatcherModel)
	//commandToHost는 호스트에 명령을 내림. 중지, 재개 등
	commandToHost(h host.HostId, cmd string)
	//requestToHost는 Watcher에 명령을 내릴 것을 Host에게 전달함
	//주로 유저의 요청을 전달
	requestToHost(w watcher.WatcherId, cmd string)

	//Inspect는 호스트의 상태를 점검함
	Inspect(hostId host.HostId) (HostInfo, HostStateData)
	//cloneHost는 입력으로 받은 호스트와 HostState가 같은 호스트를 복제함
	cloneHost(h host.Host)
}

// Server는 여러 대의 Host를 관리함
type Server struct {
	//hostsInfo는 서버가 관리중인 호스트의 외부 정보임
	hostsInfo map[host.HostId]HostInfo
	//hostsStateData는 서버가 관리중인 호스트의 내부 상태임
	hostsStateData map[host.HostId]HostStateData
	//accounts는 서버가 관리중인 계정=고객들의 정보임
	accounts map[AccountId]Account
}

// HostInfo는 호스트의 "정적 외부 정보"를 표현함
type HostInfo struct {
	hostId    host.HostId
	ipAddress IpAddress
}
type IpAddress string

// HostStateData는 호스트의 "내부 동적 상태"를 표현함
type HostStateData struct {
	// holdingWatcher는 호스트가 들고 있는 모든 Watcher리스트임.
	// 이론상 이 정보만으로 host가 들고 있는 모든 storage,poliy의 정보 파악 가능
	// 지금 단계에서 비정규화 같은 성능 이슈는 따지지 않음
	holdingWatchers []watcher.WatcherId
	performance     Performance
	machineState    host.MachineState
	watcherState    map[watcher.WatcherId]watcher.WatcherState
}

type Performance int

const (
	High Performance = iota
	Middle
	Low
	Unhealthy
)

// Account는 서비스를 위한 "계정"임
type Account struct {
	accountId      AccountId
	GlobalStorages []GlobalStorageId
	hashedPassWord AccountPassword
}
type AccountId string
type AccountPassword int

// GlobalStorageId는 "전역"에서 유일성이 유지되는 id임
// 호스트 내에선 uid만으로도 식별이 되지만,
// 여러 호스트 중에서 "고객"을 식별하려면
// hostId+userId조합이 필요
type GlobalStorageId struct {
	hid host.HostId
	sid watcher.StorageID
}
