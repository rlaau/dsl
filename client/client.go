package client

import (
	"github.com/rlaaudgjs5638/dsk/compiler"
	"github.com/rlaaudgjs5638/dsk/server"
	"github.com/rlaaudgjs5638/dsk/watcher"
)

type ClientModel interface {
	//GetNewAccount는 서버에서 계정을 부여받음
	GetNewAccount() server.Account
	// DispatchCode는 자신의 계정 정보 및 자신이 작성한 코드를 서버에 보냄
	DispatchCode(account server.Account, storageParam compiler.StorageParameter, code watcher.WatcherLanguage) error
	//CommandToWatcher는 자신이 만든 Watcher에 명령을 보냄
	CommandToWatcher(account server.Account, wid watcher.WatcherId, cmd string) error
	//InspectWatcher는 자신이 소유한 Watcher의 상태, pnl등을 점검함
	InspectWatcher(account server.Account, wid watcher.WatcherId)
}
