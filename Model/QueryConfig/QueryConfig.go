package QueryConfig

type InputFilePath struct {
	Column     string
	Data       string
	Query      string
	ValueQuery string
}

type InputData struct {
	FilePath         InputFilePath
	DataColumn       []string
	DataGrid         [][]string
	QueryString      string
	ValueQueryString string
}

type OutputData struct {
	DataColumn []string
	DataGrid   [][]string
}

type QueryConfig struct {
	TableSchema string
	Input       InputData
	Output      OutputData
}
