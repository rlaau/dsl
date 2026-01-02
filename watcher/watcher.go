package watcher

type WatcherModel interface{}

// WatcherLanguage는 Watcher언어가 accept하는 문자열 집합임
type WatcherLanguage string

// Watcher는 WatcherLang이 go에서 실행될 수 있는 형태로 변환된 결과임
type Watcher struct {
	//id는 Watcher의 식별자임
	id WatcherId
	// StorageId는 Watcher가 위임받은 자원 공간의 Id임
	storageID      StorageID
	hashedPassword StoragePassword
	//policy는 Watcher가 부여받은 행동 정책임
	policy   AST
	code     WatcherLanguage
	codeHash WatcherLangHash //받은 코드는 불변함
	//state는 현재 Watcher의 상태임
	state WatcherState
}

type WatcherId int

// Storage는 Watcher가 행동(매수,매도)할 때 쓸 수 있는 자원 저장소임
type Storage struct {
	id              StorageID
	hashedPassword  StoragePassword
	resourcesURL    []string
	apiKeys         []string
	availableAmount []int
}
type StorageID int
type StoragePassword int

type AST struct{}
type WatcherLangHash int

type WatcherState int

const (
	Mounted WatcherState = iota
	Unmounted
	Stopped
	Running
	Sleeping
)
