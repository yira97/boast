package sql

import (
	"strconv"
	"strings"
)

//PGConstraintMode is bit switch for constraint
type PGConstraintMode int64


const (
	// InvertIn means invert in to not in
	InvertIn PGConstraintMode = 1 << iota
	// LeftWhere means add where to the tap
	LeftWhere
	// SOFT means adding space to the tap and tail
	SOFT
	// InAsEqual means refactor the single in to equal
	InAsEqual
	// ZERO means no mode open
	ZERO PGConstraintMode = 0
)

// InConstraint is "in"
type InConstraint struct {
	ColumnName string
	Elected    []int64
	Mode       PGConstraintMode
	Dialect    DatabseDialect
}

func (sc *InConstraint) String() string {
	var sql strings.Builder
	if SOFT&sc.Mode > 0 {
		sql.WriteString(" ")
	}
	if len(sc.ColumnName) == 0 {
		return "ERROR_COLUMN_NAME IS EMPTY"
	}
	if len(sc.Elected) == 0 {
		return ""
	}
	if sc.Mode&LeftWhere > 0 {
		sql.WriteString("WHERE ")
	}
	sql.WriteString(sc.ColumnName)
	sql.WriteString(" ")
	if sc.Mode&InAsEqual > 0 && len(sc.Elected) == 1 {
		if sc.Mode&InvertIn > 0 {
			sql.WriteString("!= ")
			sql.WriteString(strconv.FormatInt(sc.Elected[0], 10))
		} else {
			sql.WriteString("= ")
			sql.WriteString(strconv.FormatInt(sc.Elected[0], 10))
		}
		// end. no need to space
	} else {
		if sc.Mode&InvertIn > 0 {
			sql.WriteString("NOT IN (")
		} else {
			sql.WriteString("IN (")
		}
		sql.WriteString(strconv.FormatInt(sc.Elected[0], 10))
		for i := 1; i < len(sc.Elected); i++ {
			sql.WriteString(", ")
			sql.WriteString(strconv.FormatInt(sc.Elected[i], 10))
		}
		sql.WriteString(")")
		// end. no need to space
	}

	if SOFT&sc.Mode > 0 {
		sql.WriteString(" ")
	}
	return sql.String()
}
