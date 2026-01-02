package dsl

import (
	"github.com/rlaaudgjs5638/dsk/compiler"
	"github.com/rlaaudgjs5638/dsk/host"
	"github.com/rlaaudgjs5638/dsk/server"
	"github.com/rlaaudgjs5638/dsk/view"
	"github.com/rlaaudgjs5638/dsk/watcher"
)

func MainWorkflow() {
	// 프론트엔드에서 클라이언트가 상호작용을 함
	// 클라이언트는 서버에 자동 거래 기계(=Watcher) 생성 요청 보냄.
	frontEnd := (view.ViewModel)(nil)
	clientModel := frontEnd.InteractWithClient()
	account := clientModel.GetNewAccount()
	code := watcher.WatcherLanguage("func doSomething(){}")
	storageKey := compiler.StorageParameter{ApiKey: "wqwe", AvailableAmount: 100}
	clientModel.DispatchCode(account, storageKey, code)

	// 서버는 클라이언트의 요청을 받아서, 호스트에 Watcher를 디스패치함
	server := (server.ServerModel)(nil)
	server.ReceiveCodeAndAccount(account, code)
	compiler := (compiler.CompilerModel)(nil)
	Watcher, _ := compiler.Compile(storageKey, code)
	server.DispatchWatcher(Watcher)

	//호스트는 서버로부터 Watcher를 받아서 가동함
	host := (host.HostModel)(nil)
	storagePassWord := 000_000 ///StoragePassword는 유저로부터 입력받거나, 서버 측에서 자동 생성 후 관리함
	host.MountWatcher(Watcher, watcher.StoragePassword(storagePassWord))
}
