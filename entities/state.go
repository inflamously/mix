package entities

type State struct {
	// Data must be a pointer at all times or else you will get a copy from GetState()
	Data any
}

type StateGetter[T any] interface {
	// GetState Return your own <state> type.
	GetState() *T
}
