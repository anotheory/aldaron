package Constant

var DATATYPE_WITH_QUOTE = []string{
	"character varying",
	"timestamp without time zone",
	"timestamp with time zone",
}

const INPUT_DIR = "./Input"
const OUTPUT_DIR = "./Output"

const INPUT_ORDER_DIR = "./Order/order.csv"

const COLUMN_FILE_NAME string = "column.csv"
const DATA_FILE_NAME string = "input.csv"
const QUERY_FILE_NAME string = "query.sql"
const VALUE_QUERY_FILE_NAME string = "value_query.sql"

const PLACEHOLDER_COLUMNS = "[auto:columns]"
const PLACEHOLDER_VALUES = "[auto:values]"

const REGEX_VALUES = `\[val:([0-9]+\])`

type ErrorMessage struct {
	NotFoundColumn       string
	NotFoundValue        string
	ZeroValuePlaceholder string
}

var ERROR_MESSAGE = ErrorMessage{
	NotFoundColumn:       "Input query file doesn't contain any placeholder for 'Columns'",
	NotFoundValue:        "Input query file doesn't contain any placeholder for 'Values'",
	ZeroValuePlaceholder: "Input value query file have zero 'Values' placeholder",
}
