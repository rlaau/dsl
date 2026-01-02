package dsl

type WatcherModel interface{}
type Watcher struct{}

var _ WatcherModel = (*Watcher)(nil)
