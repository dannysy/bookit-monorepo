package app

type Closable interface {
	Close() error
}

type Runnable interface {
	Start() error
	Stop() error
}
