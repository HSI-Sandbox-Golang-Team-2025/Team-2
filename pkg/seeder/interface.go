package seeder

type Seeder interface {
	RunSeeders() error
}
