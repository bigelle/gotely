package bot

// OnUpdateFunc is a handler function for processing a single incoming update.
// It receives a `Context` and returns an error if the processing fails.
type OnUpdateFunc func(Context) error

// MiddlewareFunc is a function that wraps an `OnUpdateFunc`, allowing
// pre- or post-processing of updates before passing them to the next handler.
type MiddlewareFunc func(OnUpdateFunc) OnUpdateFunc
