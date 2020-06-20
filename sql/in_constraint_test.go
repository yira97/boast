package sql

import "testing"

func TestInConstraint(t *testing.T) {
	for i, cas := range []struct {
		expect string
		cons   InConstraint
	}{
		{
			expect: "test IN (1, 2, 3)",
			cons: InConstraint{
				ColumnName: "test",
				Elected:    []int64{1, 2, 3},
				Dialect:    PG,
				Mode:       ZERO,
			},
		},
		{
			expect: "WHERE test IN (1, 2, 3)",
			cons: InConstraint{
				ColumnName: "test",
				Elected:    []int64{1, 2, 3},
				Dialect:    PG,
				Mode:       LeftWhere,
			},
		},
		{
			expect: "WHERE test NOT IN (1, 2, 3)",
			cons: InConstraint{
				ColumnName: "test",
				Elected:    []int64{1, 2, 3},
				Dialect:    PG,
				Mode:       LeftWhere | InvertIn,
			},
		},
		{
			expect: " WHERE test NOT IN (1, 2, 3) ",
			cons: InConstraint{
				ColumnName: "test",
				Elected:    []int64{1, 2, 3},
				Dialect:    PG,
				Mode:       LeftWhere | InvertIn | SOFT,
			},
		},
		{
			expect: " WHERE test != 1 ",
			cons: InConstraint{
				ColumnName: "test",
				Elected:    []int64{1},
				Dialect:    PG,
				Mode:       LeftWhere | InvertIn | SOFT | InAsEqual,
			},
		},
	} {
		if cas.cons.String() != cas.expect {
			t.Error("in constraint fail at index", i)
		}
	}
}
