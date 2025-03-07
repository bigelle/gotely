package bot

// OnUpdateFunc is a function that is used to respond to one incoming update
type OnUpdateFunc func(Context) error

type MiddlewareFunc func(OnUpdateFunc) OnUpdateFunc
