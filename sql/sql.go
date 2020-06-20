package sql

// DatabseDialect is to differenciate database syntax for each database
type DatabseDialect int

const (
	// PG is postgreSQL
	PG DatabseDialect = 1 << iota
)
