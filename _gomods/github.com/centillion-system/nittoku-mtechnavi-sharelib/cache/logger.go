package cache

type Logger interface {
	WithError(error) Logger

	Warnf(string, ...interface{})
}
