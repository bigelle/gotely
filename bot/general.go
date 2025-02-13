package bot

// OnUpdateFunc is a function that is used to respond to one incoming update
type OnUpdateFunc func(Context) error

// MiddlewareFunc is a function that is used before responding to every update
type MiddlewareFunc func(Context) error
