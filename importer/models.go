package importer

type Workday struct {
	Date     string `csv:"Data"`
	In       string `csv:"Entrada Empresa"`
	BreakOut string `csv:"Saida Intervalo"`
	BreakIn  string `csv:"Entrada Intervalo"`
	Out      string `csv:"Saida Empresa"`
	ExtraIn  string `csv:"Entrada Extra"`
	ExtraOut string `csv:"Saida Extra"`
	Balance  string `csv:"Saldo"`
}
