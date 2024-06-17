package dip

type DataService struct {
	ReadData DataReader
}

func (s DataService) GetData() string {
	return s.ReadData.Read()
}
