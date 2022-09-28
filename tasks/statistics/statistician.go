package statistics

type Statistician interface {
	GetStatistics() (*TextStatistics, error)
}
