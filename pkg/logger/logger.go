package logger

// Logger ...
type Logger interface {
	Log(source string, event string, data interface{}) error
	LogError(source string, event string, data interface{}) error
	LogWarning(source string, event string, data interface{}) error
}
