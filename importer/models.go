package importer

type test struct {
	Name string  `csv:"header1"`
	ID   int     `csv:"header2"`
	Num  float64 `csv:"header3"`
}
