package types

type MalType struct {
	intVal int
	DataType string
	stringVal string
}

type MalNumber struct {
	MalType // anonymous field
}

type MalSymbol struct {
	MalType // anonymous field
}
