package store

var factory Factory

type Factory interface {
	ApiStore() ApiStore
	Close() error
}

func GetFactory() Factory {
	return factory
}

func SetFactory(f Factory) {
	factory = f
}
