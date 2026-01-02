package compiler

import "github.com/rlaaudgjs5638/dsk/watcher"

type CompilerModel interface {
	//Compile은 스토리지를 만들기 위한 매게변수와 코드를 받아서 Watcher로 변환함
	Compile(storageParameter StorageParameter, code watcher.WatcherLanguage) (watcher.Watcher, error)
	//makeWatcherId는 Compile과정에서 Watcher에 WatcherId를 부여함. id는 Storage와 policy의
	makeWatcherId(sid watcher.StorageID, code watcher.WatcherLanguage) watcher.WatcherId //hash로 id생성
}

// BuildStorage는  StorageParameter기반으로 Watcher를 위한 Storage를 만듦.
func BuildStorage(sp StorageParameter) watcher.Storage

type StorageParameter struct {
	ApiKey          string
	AvailableAmount int
}
