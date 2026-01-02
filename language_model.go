package dsl

type SourceCodeModel interface {
	Run() error
	Analyze() Watcher
}
