package repositories

type ApplicationsRepository interface {
	GetApplicationNames() ([]string, error)
}
