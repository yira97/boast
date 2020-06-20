package utils

import (
	"testing"
)

func TestPPToLF(t *testing.T) {
	for i, cas := range []struct {
		page         int
		pageSzie     int
		expectLimit  int
		expectOffset int
	}{
		{1, 20, 20, 0},
		{0, 20, 20, 0},
		{2, 10, 10, 10},
	} {
		limit, offset := PPToLF(cas.page, cas.pageSzie)
		if limit != cas.expectLimit {
			t.Error("limit error. index", i)
		}
		if offset != cas.expectOffset {
			t.Error("offset error index", i)
		}
	}
}
