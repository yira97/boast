package sql

// Constraint can construct where clause
type Constraint interface {
	String() string
	Merge(...Constraint)
}
