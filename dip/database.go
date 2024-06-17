package dip

type Database struct{}

func (d Database) Read() string {
	return "Reading data from database..."
}
