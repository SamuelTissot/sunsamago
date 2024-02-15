package app

type Stopper interface {
	Done() <-chan struct{}
}
