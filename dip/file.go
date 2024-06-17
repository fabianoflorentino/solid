package dip

type File struct{}

func (f File) Read() string {
	return "Reading data from file..."
}
