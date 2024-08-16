package entities

type Initializer interface {
	Initialize(game GameRunner) error
}
