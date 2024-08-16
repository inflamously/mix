package entities

type Updater interface {
	Update(game GameRunner) error
}
